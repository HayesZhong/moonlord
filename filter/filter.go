package filter

import ()

type Filter interface {
	Filter() (fx, fy []float64, err error)
}
