package ica

import (
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

type ICA struct {
	xMat *mat.Dense
}

func NewICA(x [][]float64) *ICA {
	ica := new(ICA)
	x1d := make([]float64, len(x)*len(x[0]))
	for i, inSl := range x {
		for j, v := range inSl {
			x1d[len(x)*i+j] = v
		}
	}
	ica.xMat = mat.NewDense(len(x), len(x[0]), x1d)
	return ica
}

func (ica *ICA) CalcICA() [][]float64 {
	ica.fit()
	z := ica.whiten()
	y := ica.analyze(z)
	return y
}

func (ica *ICA) fit() {
	r, c := ica.xMat.Dims()
	sig := make([]float64, len(r))
	for i := 0; i < c; i++ {
		mat.Col(sig, i, ica.xMat)
		floats.Sub(sig, floats.)
	}
}
