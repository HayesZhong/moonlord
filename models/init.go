package models

import (
	"moonlord/filter"
)

var MTras [][]Tra
var RTrees map[int]*Rtree
var minChildren int

func init() {
	mtrasNum := 20000
	minChildren = int(mtrasNum / 100)
	traDatas, _ := GetTraData("F:/tradata/format", mtrasNum+1)
	MTras = traDatas
	FilteMultiTras(MTras)
	RTrees = make(map[int]*Rtree)
	go func() {
		for _, tras := range MTras {
			InsterTrasToRTrees(RTrees, tras)
		}
	}()
}

func InsterTrasToRTrees(rTrees map[int]*Rtree, tras []Tra) {
	if rTrees[tras[0].Zone] == nil {
		rTrees[tras[0].Zone] = NewTree(3, minChildren, minChildren*2)
	}
	minx, miny, mint, maxx, maxy, maxt := GetTrasMinAndMax(tras)
	rTrees[tras[0].Zone].Insert(NewRectWithTwoPoint([]float64{minx, miny, mint}, []float64{maxx, maxy, maxt}, tras))
}

func GetTrasMinAndMax(tras []Tra) (minx, miny, mint, maxx, maxy, maxt float64) {
	minx, miny, mint = tras[0].FX, tras[0].FY, float64(tras[0].Time.Unix())
	maxx, maxy, maxt = tras[0].FX, tras[0].FY, float64(tras[0].Time.Unix())
	for i := 1; i < len(tras); i++ {
		if minx > tras[i].FX {
			minx = tras[i].FX
		}
		if miny > tras[i].FY {
			miny = tras[i].FY
		}
		if mint > float64(tras[i].Time.Unix()) {
			mint = float64(tras[i].Time.Unix())
		}
		if maxx < tras[i].FX {
			maxx = tras[i].FX
		}
		if maxy < tras[i].FY {
			maxy = tras[i].FY
		}
		if maxt < float64(tras[i].Time.Unix()) {
			maxt = float64(tras[i].Time.Unix())
		}
	}
	return
}

func FilteMultiTras(mtras [][]Tra) {
	k := new(filter.KlamanFilter)
	for i := 0; i < len(mtras); i++ {
		FilteOneTras(mtras[i], k)
	}
}

func FilteOneTras(tras []Tra, k *filter.KlamanFilter) {
	x := make([]float64, len(tras))
	y := make([]float64, len(tras))
	t := make([]float64, len(tras))
	tras[0].LatLon2XY()
	x[0] = 0
	y[0] = 0
	t[0] = 0
	for j := 1; j < len(tras); j++ {
		tras[j].LatLon2XY()
		x[j] = tras[j].X - tras[0].X
		y[j] = tras[j].Y - tras[0].Y
		t[j] = float64(tras[j].Time.Unix() - tras[0].Time.Unix())
	}

	k.X, k.Y, k.T, k.N = x, y, t, len(x)
	fx, fy := k.Filter()
	tras[0].FX = tras[0].X
	tras[0].FY = tras[0].Y
	tras[0].XY2LatLon()
	for j := 1; j < len(tras); j++ {
		tras[j].FX = tras[0].X + fx[j]
		tras[j].FY = tras[0].Y + fy[j]
		tras[j].XY2LatLon()
	}
}
