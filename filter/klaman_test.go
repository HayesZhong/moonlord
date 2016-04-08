package filter

import ()

//func TestKlamanFilter(t *testing.T) {
//	k := KlamanFilter{
//		X: []float64{1, 2, 3},
//		Y: []float64{1, 2, 3},
//		T: []float64{1, 2, 3},
//	}
//	x, y := k.Filter()
//	if !(x[0] == 1 && x[1] == 2 && x[2] == 3 && y[0] == 1 && y[1] == 2 && y[2] == 3) {
//		t.Fail()
//	}
//}

//func TestKlamanFilter2(t *testing.T) {
//	index := 3
//	traDatas, err := models.GetTraData("F:/tradata/format", index+1)
//	if err != nil {
//		t.Fail()
//	}
//	for j := 0; j < len(traDatas); j++ {
//		for i := 0; i < len(traDatas[j]); i++ {
//			traDatas[j][i].LatLon2XY()
//		}
//	}
//	x := make([]float64, len(traDatas[index]))
//	y := make([]float64, len(traDatas[index]))
//	tt := make([]float64, len(traDatas[index]))
//	x[0] = 0
//	y[0] = 0
//	tt[0] = 0
//	for i := 1; i < len(traDatas[index]); i++ {
//		x[i] = (float64(traDatas[index][i].X) - float64(traDatas[index][0].X))
//		y[i] = (float64(traDatas[index][i].Y) - float64(traDatas[index][0].Y))
//		tt[i] = float64(traDatas[index][i].Time.UnixNano()) / float64(10e8)
//	}
//	f := KlamanFilter{
//		X: x,
//		Y: y,
//		T: tt,
//	}
//	fx, fy, _ := f.Filter()
//	t.Log(x)
//	t.Log(fx)
//	t.Log(y)
//	t.Log(fy)
//}
