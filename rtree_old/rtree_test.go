package rtree_old

import "testing"

type P struct {
	I int
}

func Benchmark_insertNearest(b *testing.B) {
	//b2 := time.Now().UnixNano() / 10e6
	//	ps := make([]P, 1000000)
	//	things := make([]*rtree.Rect, 1000000)
	//	rt := rtree.NewTree(2, 3, 3)
	//	for i := 0; i < 1000000; i++ {
	//		ps[i] = P{i}
	//		things[i] = rtree.NewRect(rtree.NewPointWithData([]float64{float64(rand.Intn(1000000)), float64(rand.Intn(1000000))}, &ps[i]), []float64{float64(rand.Intn(1000)), float64(rand.Intn(1000))})
	//		rt.Insert(things[i])
	//	}
	//	a2 := time.Now().UnixNano() / 10e6
	//
	//	fmt.Println(a2 - b2)
	//
	//	b := time.Now().UnixNano() / 10e6
	//
	//	objs := rt.NearestNeighbors(3, *rtree.NewPointWithData([]float64{float64(rand.Intn(1000000)), float64(rand.Intn(1000000))}, nil))
	//	a := time.Now().UnixNano() / 10e6
	//	fmt.Println(a - b)
	//	o, bl := objs[2].Bounds().GetP().Obj.(P)
	//	if bl {
	//		fmt.Println(o.I)
	//	}
}
