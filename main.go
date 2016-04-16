package main

import (
	_ "moonlord/models"
	_ "moonlord/routers"
	"runtime"

	"github.com/astaxie/beego"
)

func main() {
	threadNum := runtime.NumCPU()
	runtime.GOMAXPROCS(threadNum)
	beego.Run()
}

//package main

//import (
//	"fmt"
//	"math"
//)

//type Tra struct {
//	Lat  float64 `json:"lat"`
//	Lon  float64 `json:"lon"`
//	Time float64 `json:"time"`
//	X    float64 `json:"x"`
//	Y    float64 `json:"y"`
//	FLat float64 `json:"flat"`
//	FLon float64 `json:"flon"`
//	FX   float64 `json:"fx"`
//	FY   float64 `json:"fy"`
//}

//type VMatrix struct {
//	Vx float64
//	Vy float64
//}

//type XMatrix struct {
//	X float64
//	Y float64
//}

//func main() {
//	S := make([]Tra, 10)
//	R := make([]Tra, 11)

//	fmt.Printf("%v", len(S))
//	for i := 0; i < 10; i++ {
//		R[i+1].FX = float64(i)
//		R[i+1].FY = float64(i)
//		R[i+1].Time = float64(i)

//		S[i].FX = float64(i)
//		S[i].FY = float64(i)
//		S[i].Time = float64(i)
//	}
//	a := ERPDistance(S, R, Tra{X: 0, Y: 0})
//	//a := DTWDistance(S, R, 0, 0)
//	fmt.Printf("%f", a)
//}

//func ERPDistance(S []Tra, R []Tra, randomPoint Tra) float64 {
//	var simDis float64
//	//S is empty or R is empty
//	if len(S) == 0 {
//		for _, v := range R {
//			simDis = simDis + EuclideanDistance(randomPoint, v)
//		}
//		return simDis
//	} else {
//		if len(R) == 0 {
//			for _, v := range S {
//				simDis = simDis + EuclideanDistance(randomPoint, v)
//			}
//			return simDis
//		}
//	}

//	//this is recursion part
//	var min1 float64 = math.Min(ERPDistance(S[1:], R[1:], randomPoint)+EuclideanDistance(S[0], R[0]), ERPDistance(S[1:], R, randomPoint)+EuclideanDistance(S[0], randomPoint))
//	return math.Min(min1, ERPDistance(S, R[1:], randomPoint)+EuclideanDistance(R[0], randomPoint))
//}

//func DTWDistance(S []Tra, R []Tra, Is int, Ir int) float64 {
//	if len(S) == Is+1 && len(R) == Ir+1 {
//		return 0
//	} else {
//		if len(S) == Is+1 || len(R) == Ir+1 {
//			//此处得给该矩形块中的对角线最大值
//			return 9999
//		}
//	}

//	min1 := math.Min(DTWDistance(S, R, Is, Ir+1), DTWDistance(S, R, Is+1, Ir))
//	min := TightenedDistance(S, R, Is, Ir) + math.Min(min1, DTWDistance(S, R, Is+1, Ir+1))
//	return min

//}

//func EuclideanDistance(A Tra, B Tra) float64 {
//	return math.Sqrt(A.FX-B.FX)*(A.FX-B.FX) + (A.FY-B.FY)*(A.FY-B.FY)
//}

//func TightenedDistance(S []Tra, R []Tra, Is int, Ir int) float64 {
//	//计算S[I]的速度向量
//	U := VMatrix{(S[Is].FX - S[Is+1].FX) / (S[Is].Time - S[Is+1].Time), (S[Is].FY - S[Is+1].FY) / (S[Is].Time - S[Is+1].Time)}
//	//计算R[I]的速度向量
//	V := VMatrix{(R[Ir].FX - R[Ir+1].FX) / (R[Ir].Time - R[Ir+1].Time), (R[Ir].FY - R[Ir+1].FY) / (R[Ir].Time - R[Ir+1].Time)}
//	UV := VMatrix{math.Abs(U.Vx - V.Vx), math.Abs(U.Vy - V.Vy)}
//	//计算Tcpa
//	Tcpa := (math.Abs(S[Is].FX-R[Ir].FX)*UV.Vx + math.Abs(S[Is].FY-R[Ir].FY)*UV.Vy) / (UV.Vx*UV.Vx + UV.Vy*UV.Vy)
//	//重叠时间下限
//	Tlow := math.Max(S[Is].Time, R[Ir].Time)
//	//重叠时间上限
//	Tup := math.Min(S[Is+1].Time, R[Ir+1].Time)

//	if Tlow >= Tup {
//		//当前矩形块的最大值
//		return 9999
//	}

//	if Tcpa < Tlow {
//		//如果Tcpa在重叠时间的左侧
//		Tcpa = Tlow
//	} else {
//		if Tcpa > Tup {
//			//如果Tcpa在重叠时间的右侧
//			Tcpa = Tup
//		}
//	}

//	if U.Vx == V.Vx && U.Vy == V.Vy {
//		Tcpa = Tlow
//	}

//	//生成Tcpa时间S轨迹和R轨迹最近的点
//	Sx := XMatrix{S[Is].FX + U.Vx*Tcpa, S[Is].FY + U.Vy*Tcpa}
//	Rx := XMatrix{R[Ir].FX + V.Vx*Tcpa, R[Ir].FY + V.Vy*Tcpa}

//	return math.Sqrt((Sx.X-Rx.X)*(Sx.X-Rx.X) + (Sx.Y-Rx.Y)*(Sx.Y-Rx.Y))
//}
