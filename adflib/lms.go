package adflib

import (
	"errors"
	"github.com/gonum/floats"
)

type FiltLMS struct {
	AdaptiveFilter
	kind string
	wHistory [][]float64
}

func NewFiltLMS(n int, mu float64, w interface{}) (*FiltLMS, error) {
	var err error
	p := new(FiltLMS)
	p.kind = "LMS filter"
	p.n = n
	p.mu, err = p.CheckFloatParam(mu, 0, 1000, "mu")
	if err != nil {
		return nil, err
	}
	err = p.InitWeghts(w, n)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (af *FiltLMS) adapt(d float64, x []float64) {
	y := floats.Dot(af.w, x)
	e := d - y
	for i := 0; i < len(x); i++ {
		af.w[i] += af.mu * e * x[i]
	}
}

func (af *FiltLMS) Run(d []float64, x [][]float64) ([]float64, []float64, [][]float64, error) {
	//measure the data and check if the dimmension agree
	N := len(x)
	if len(d) != N {
		return nil, nil, nil, errors.New("The length of slice d and x must agree.")
	}
	af.n = len(x[0])
	af.wHistory = make([][]float64, af.n)

	y := make([]float64, N)
	e := make([]float64, N)
	//adaptation loop
	for i:=0;i< N;i++{
		af.wHistory[i] = d
		y[i] = floats.Dot(af.w, x[i])
		e[i] = d[i] - y[i]
		for j:=0;j<af.n;j++{
			af.w[j] = af.mu * e[i] * x[i][j]
		}
	}
	return y, e, af.wHistory, nil
}
