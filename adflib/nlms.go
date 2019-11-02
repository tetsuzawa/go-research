package adflib

import (
	"errors"

	"github.com/gonum/floats"
)

type FiltNLMS struct {
	AdaptiveFilter
	kind     string
	eps      float64
	wHistory [][]float64
}

func NewFiltNLMS(n int, mu float64, eps float64, w interface{}) (*FiltNLMS, error) {
	var err error
	p := new(FiltNLMS)
	p.kind = "NLMS filter"
	p.n = n
	p.mu, err = p.CheckFloatParam(mu, 0, 1000, "mu")
	if err != nil {
		return nil, err
	}
	p.eps, err = p.CheckFloatParam(eps, 0, 1000, "eps")
	if err != nil {
		return nil, err
	}
	err = p.InitWeights(w, n)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (af *FiltNLMS) Adapt(d float64, x []float64) {
	y := floats.Dot(af.w, x)
	e := d - y
	nu := af.mu / (af.eps + floats.Dot(x, x))
	for i := 0; i < len(x); i++ {
		af.w[i] += nu * e * x[i]
	}
}

func (af *FiltNLMS) Run(d []float64, x [][]float64) ([]float64, []float64, [][]float64, error) {
	//measure the data and check if the dimension agree
	N := len(x)
	if len(d) != N {
		return nil, nil, nil, errors.New("The length of slice d and x must agree.")
	}
	af.n = len(x[0])
	af.wHistory = make([][]float64, N)

	y := make([]float64, N)
	e := make([]float64, N)
	//adaptation loop
	for i := 0; i < N; i++ {
		af.wHistory[i] = af.w
		y[i] = floats.Dot(af.w, x[i])
		e[i] = d[i] - y[i]
		nu := af.mu / (af.eps + floats.Dot(x[i], x[i]))
		for j := 0; j < af.n; j++ {
			af.w[j] = nu * e[i] * x[i][j]
		}
	}
	return y, e, af.wHistory, nil
}
