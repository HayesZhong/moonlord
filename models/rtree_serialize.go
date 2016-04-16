package models

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
)

var tmp_int64 int64 = 0
var tmp_int8 int8 = 0
var tmp_uint64 uint64 = 0

func (tree *Rtree) Encode(filePath string) error {
	file, e := os.Create(filePath)
	if e != nil {
		fmt.Println(e)
		return e
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	binary.Write(writer, binary.BigEndian, int64(tree.Dim))
	binary.Write(writer, binary.BigEndian, int64(tree.MinChildren))
	binary.Write(writer, binary.BigEndian, int64(tree.MaxChildren))
	binary.Write(writer, binary.BigEndian, int64(tree.Size))
	binary.Write(writer, binary.BigEndian, int64(tree.Height))
	tree.Root.Encode(writer)
	writer.Flush()
	return nil
}

func (tree *Rtree) Decode(filePath string) error {
	file, e := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if e != nil {
		fmt.Println(e)
		return e
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	binary.Read(reader, binary.BigEndian, &tmp_int64)
	tree.Dim = int(tmp_int64)
	binary.Read(reader, binary.BigEndian, &tmp_int64)
	tree.MinChildren = int(tmp_int64)
	binary.Read(reader, binary.BigEndian, &tmp_int64)
	tree.MaxChildren = int(tmp_int64)
	binary.Read(reader, binary.BigEndian, &tmp_int64)
	tree.Size = int(tmp_int64)
	binary.Read(reader, binary.BigEndian, &tmp_int64)
	tree.Height = int(tmp_int64)

	tree.Root = &node{}
	tree.Root.Decode(reader)
	setParent(tree.Root)
	return nil
}

func setParent(node *node) {
	if node == nil || node.Leaf == true {
		return
	}
	for _, entry := range node.Entries {
		entry.Child.parent = node
		setParent(entry.Child)
	}
}

func (node *node) Encode(writer io.Writer) {
	if node.Leaf {
		binary.Write(writer, binary.BigEndian, int8(1))
	} else {
		binary.Write(writer, binary.BigEndian, int8(0))
	}
	binary.Write(writer, binary.BigEndian, int64(node.Level))
	binary.Write(writer, binary.BigEndian, int64(len(node.Entries)))
	for i := 0; i < len(node.Entries); i++ {
		node.Entries[i].Encode(writer)
	}
}

func (node *node) Decode(reader io.Reader) {
	binary.Read(reader, binary.BigEndian, &tmp_int8)
	if tmp_int8 == 1 {
		node.Leaf = true
	} else {
		node.Leaf = false
	}
	binary.Read(reader, binary.BigEndian, &tmp_int64)
	node.Level = int(tmp_int64)

	binary.Read(reader, binary.BigEndian, &tmp_int64)
	l := int(tmp_int64)
	node.Entries = make([]entry, l, l)
	for i := 0; i < l; i++ {
		node.Entries[i] = entry{}
		node.Entries[i].Decode(reader)
	}

}

func (entry *entry) Encode(writer io.Writer) {
	entry.Bb.Encode(writer)
	if entry.Child != nil {
		binary.Write(writer, binary.BigEndian, int8(1))
		entry.Child.Encode(writer)
	} else {
		binary.Write(writer, binary.BigEndian, int8(0))
	}

}

func (entry *entry) Decode(reader io.Reader) {
	entry.Bb = &Rect{}
	entry.Bb.Decode(reader)
	entry.Obj = entry.Bb

	binary.Read(reader, binary.BigEndian, &tmp_int8)
	if tmp_int8 == 1 {
		entry.Child = &node{}
		entry.Child.Decode(reader)
	} else {
	}
}

func (rect *Rect) Encode(writer io.Writer) {
	binary.Write(writer, binary.BigEndian, []uint64{math.Float64bits(float64(rect.P[0])), math.Float64bits(float64(rect.P[1])), math.Float64bits(float64(rect.P[2]))})
	binary.Write(writer, binary.BigEndian, []uint64{math.Float64bits(float64(rect.Q[0])), math.Float64bits(float64(rect.Q[1])), math.Float64bits(float64(rect.Q[2]))})
	binary.Write(writer, binary.BigEndian, int64(len(rect.Tras)))
	for i := 0; i < len(rect.Tras); i++ {
		rect.Tras[i].Encode(writer)
	}

}

func (rect *Rect) Decode(reader io.Reader) {
	rect.P = make([]float64, 3, 3)
	for i := 0; i < 3; i++ {
		binary.Read(reader, binary.BigEndian, &tmp_uint64)
		rect.P[i] = math.Float64frombits(tmp_uint64)
	}
	rect.Q = make([]float64, 3, 3)
	for i := 0; i < 3; i++ {
		binary.Read(reader, binary.BigEndian, &tmp_uint64)
		rect.Q[i] = math.Float64frombits(tmp_uint64)
	}
	binary.Read(reader, binary.BigEndian, &tmp_int64)
	l := int(tmp_int64)
	rect.Tras = make([]Tra, l, l)
	for i := 0; i < l; i++ {
		rect.Tras[i] = Tra{}
		rect.Tras[i].Decode(reader)
	}
}

func (tra *Tra) Encode(writer io.Writer) {
	binary.Write(writer, binary.BigEndian, math.Float64bits(tra.X))
	binary.Write(writer, binary.BigEndian, math.Float64bits(tra.Y))
	binary.Write(writer, binary.BigEndian, math.Float64bits(tra.FX))
	binary.Write(writer, binary.BigEndian, math.Float64bits(tra.FY))
	binary.Write(writer, binary.BigEndian, math.Float64bits(tra.FLat))
	binary.Write(writer, binary.BigEndian, math.Float64bits(tra.FLon))
	binary.Write(writer, binary.BigEndian, math.Float64bits(tra.Lat))
	binary.Write(writer, binary.BigEndian, math.Float64bits(tra.Lon))
	binary.Write(writer, binary.BigEndian, tra.T)
}
func (tra *Tra) Decode(reader io.Reader) {
	binary.Read(reader, binary.BigEndian, &tmp_uint64)
	tra.X = math.Float64frombits(tmp_uint64)
	binary.Read(reader, binary.BigEndian, &tmp_uint64)
	tra.Y = math.Float64frombits(tmp_uint64)
	binary.Read(reader, binary.BigEndian, &tmp_uint64)
	tra.FX = math.Float64frombits(tmp_uint64)
	binary.Read(reader, binary.BigEndian, &tmp_uint64)
	tra.FY = math.Float64frombits(tmp_uint64)
	binary.Read(reader, binary.BigEndian, &tmp_uint64)
	tra.FLat = math.Float64frombits(tmp_uint64)
	binary.Read(reader, binary.BigEndian, &tmp_uint64)
	tra.FLon = math.Float64frombits(tmp_uint64)
	binary.Read(reader, binary.BigEndian, &tmp_uint64)
	tra.Lat = math.Float64frombits(tmp_uint64)
	binary.Read(reader, binary.BigEndian, &tmp_uint64)
	tra.Lon = math.Float64frombits(tmp_uint64)

	binary.Read(reader, binary.BigEndian, &tmp_int64)
	tra.T = tmp_int64
}

//func (tree *Rtree) Encode(filePath string) error {

//	file, e := os.Create(filePath)
//	if e != nil {
//		fmt.Println(e)
//		return e
//	}
//	defer file.Close()
//	gob.Register(&Rect{})
//	enc := gob.NewEncoder(file)
//	err := enc.Encode(*tree)
//	if err != nil {
//		return err
//	}

//	return nil
//}

//func (tree *Rtree) Decode(filePath string) error {

//	file, e := os.OpenFile(filePath, os.O_RDONLY, 0666)
//	if e != nil {
//		return e
//	}
//	defer file.Close()

//	gob.Register(&Rect{})
//	enc := gob.NewDecoder(file)
//	err := enc.Decode(tree)
//	if err != nil {
//		return err
//	}
//	node := tree.Root
//	setParent(node)

//	return nil
//}

//func setParent(node *node) {
//	if node == nil || node.Leaf == true {
//		return
//	}
//	for _, entry := range node.Entries {
//		entry.Child.parent = node
//		setParent(entry.Child)
//	}
//}
