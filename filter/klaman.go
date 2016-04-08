package filter

import (
	"trasim/matrix"
)

type KlamanFilter struct {
	X    []float64
	Y    []float64
	T    []float64
	N    int //长度
	xn   [][][]float64
	xn_1 [][][]float64
	kn   [][][]float64
	pn   [][][]float64
	pn_1 [][][]float64
}

const (
	SIGMA   = 4
	SIGMA_S = 1
)

var Q, H, R, I [][]float64

func initQHRI() {
	Q = make([][]float64, 4, 4)
	for i := 0; i < 4; i++ {
		Q[i] = make([]float64, 4, 4)
	}
	Q[2][2] = SIGMA_S * SIGMA_S
	Q[3][3] = SIGMA_S * SIGMA_S

	H = make([][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		H[i] = make([]float64, 4, 4)
	}
	H[0][0] = 1
	H[1][1] = 1

	R = make([][]float64, 2, 2)
	for i := 0; i < 2; i++ {
		R[i] = make([]float64, 2, 2)
		R[i][i] = SIGMA * SIGMA
	}

	I = make([][]float64, 4, 4)
	for i := 0; i < 4; i++ {
		I[i] = make([]float64, 4, 4)
		I[i][i] = 1
	}
}

func (k *KlamanFilter) init() {
	k.xn = make([][][]float64, k.N)
	k.kn = make([][][]float64, k.N)
	k.xn_1 = make([][][]float64, k.N)
	k.pn = make([][][]float64, k.N)
	k.pn_1 = make([][][]float64, k.N)
}

func (k *KlamanFilter) Filter() (fx, fy []float64) {
	k.init()
	k.xnf(k.N - 1)
	fx = make([]float64, k.N)
	fy = make([]float64, k.N)
	for i := 0; i < k.N; i++ {
		fx[i] = k.xn[i][0][0]
		fy[i] = k.xn[i][1][0]
	}
	return fx, fy
}

func (k *KlamanFilter) getA(i int) [][]float64 {
	r := make([][]float64, 4, 4)
	for j := 0; j < 4; j++ {
		r[j] = make([]float64, 4, 4)
		r[j][j] = 1
	}
	r[0][2] = k.T[i] - k.T[i-1]
	r[1][3] = k.T[i] - k.T[i-1]
	return r
}
func (k *KlamanFilter) xnf(i int) [][]float64 {
	if k.xn[i] != nil {
		return k.xn[i]
	}
	if i == 0 {
		r := make([][]float64, 4)
		for ii := 0; ii < 4; ii++ {
			r[ii] = make([]float64, 1)
		}
		r[0][0] = k.X[0]
		r[1][0] = k.Y[0]
		r[2][0] = (k.X[1] - k.X[0]) / (k.T[1] - k.T[0])
		r[3][0] = (k.Y[1] - k.Y[0]) / (k.T[1] - k.T[0])
		k.xn[0] = r
		return r
	}
	xn_1f := k.xn_1f(i)
	t1 := xn_1f
	t2 := matrix.DotProduct(H, xn_1f)
	t3 := matrix.Subtract([][]float64{{k.X[i]}, {k.Y[i]}}, t2)
	t4 := matrix.DotProduct(k.knf(i), t3)
	r := matrix.Add(t1, t4)
	k.xn[i] = r
	return r
}

func (k *KlamanFilter) xn_1f(i int) [][]float64 {
	if i <= 0 {
		return nil
	}
	if k.xn_1[i] != nil {
		return k.xn_1[i]
	}
	r := matrix.DotProduct(k.getA(i), k.xnf(i-1))
	k.xn_1[i] = r
	return r
}

func (k *KlamanFilter) pnf(i int) [][]float64 {
	if k.pn[i] != nil {
		return k.pn[i]
	}
	if i == 0 {
		r := make([][]float64, 4, 4)
		for i := 0; i < 4; i++ {
			r[i] = make([]float64, 4, 4)
			r[i][i] = SIGMA_S * SIGMA_S
		}
		k.pn[0] = r
		return r
	}
	pn_1f := k.pn_1f(i)
	t1 := matrix.DotProduct(k.knf(i), H)
	t2 := matrix.Subtract(I, t1)
	r := matrix.DotProduct(t2, pn_1f)
	k.pn[i] = r
	return r

}

func (k *KlamanFilter) pn_1f(i int) [][]float64 {
	if k.pn_1[i] != nil {
		return k.pn_1[i]
	}
	if i <= 0 {
		return nil
	}
	A := k.getA(i)
	t1 := matrix.DotProduct(A, k.pnf(i-1))
	t2 := matrix.DotProduct(t1, matrix.Transpose(A))
	r := matrix.Add(t2, Q)
	k.pn_1[i] = r
	return r
}

func (k *KlamanFilter) knf(i int) [][]float64 {
	if k.kn[i] != nil {
		return k.kn[i]
	}
	pn_1f := k.pn_1f(i)
	t1 := matrix.DotProduct(pn_1f, matrix.Transpose(H))
	t2 := matrix.DotProduct(matrix.DotProduct(H, pn_1f), matrix.Transpose(H))
	r := matrix.DotProduct(t1, matrix.Inverse(matrix.Add(R, t2)))
	k.kn[i] = r
	return r
}
