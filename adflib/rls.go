package adflib

import (
	"gonum.org/v1/gonum/mat"
)

type FiltRLS struct {
	AdaptiveFilter
	kind     string
	wHistory *mat.Dense
	eps      float64
	R        *mat.Dense
}

func NewFiltRLS(n int, mu float64, eps float64, w interface{}) (ADFInterface, error) {
	var err error
	p := new(FiltRLS)
	p.kind = "RLS filter"
	p.n = n
	p.mu, err = p.CheckFloatParam(mu, 0, 1, "mu")
	if err != nil {
		return nil, err
	}
	p.eps, err = p.CheckFloatParam(mu, 0, 1, "eps")
	if err != nil {
		return nil, err
	}
	err = p.InitWeights(w, n)
	if err != nil {
		return nil, err
	}
	var Rs = make([]float64, n*n)
	for i := 0; i < n; i++ {
		Rs[i*(n+1)] = 1 / eps
	}
	p.R = mat.NewDense(n, n, Rs)
	return p, nil
}

/*
//TODO baseをgonumに書き換えてから書く
func (af *FiltRLS) Adapt(d float64, x []float64) {
		y := floats.Dot(af.w, x)
		e := d - y
		for i := 0; i < len(x); i++ {
			af.w[i] += af.mu * e * x[i]
		}
}

func (af *FiltLMS) Run(d []float64, x [][]float64) ([]float64, []float64, [][]float64, error) {
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
		for j := 0; j < af.n; j++ {
			af.w[j] = af.mu * e[i] * x[i][j]
		}
	}
	return y, e, af.wHistory, nil
}
*/
