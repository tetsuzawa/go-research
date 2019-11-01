package adflib

import (
	"errors"
	"github.com/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

type FiltAP struct {
	AdaptiveFilter
	kind     string
	order    int
	eps      float64
	wHistory *mat.Dense
	xMem     *mat.Dense
	dMem     *mat.Dense
	yMem     *mat.Dense
	eMem     *mat.Dense
	epsIDE   *mat.Dense
	ide      *mat.Dense
}

func NewAP(n int, mu float64, order int, eps float64, w interface{}) (*FiltAP, error) {
	var err error
	p := new(FiltAP)
	p.kind = "AP filter"
	p.n = n
	p.mu, err = p.CheckFloatParam(mu, 0, 1000, "mu")
	if err != nil {
		return nil, err
	}
	p.order = order
	p.eps, err = p.CheckFloatParam(eps, 0, 1000, "eps")
	if err != nil {
		return nil, err
	}
	err = p.InitWeights(w, n)
	if err != nil {
		return nil, err
	}
	p.xMem = mat.NewDense(n, order, nil)
	p.dMem = mat.NewDense(1, order, nil)

	elmNum := order * order

	//make diagonal matrix
	diaMat := make([]float64, elmNum)
	for i := 0; i < order; i++ {
		diaMat[i*(order+1)] = eps
	}
	p.epsIDE = mat.NewDense(order, order, diaMat)

	for i := 0; i < order; i++ {
		diaMat[i*(order+1)] = 1
	}
	p.ide = mat.NewDense(order, order, diaMat)
	p.wHistory = mat.NewDense(n, order, nil)

	return p, nil
}

func (af *FiltAP) Adapt(d float64, x []float64) error {
	xr, _ := af.xMem.Dims()
	xCol := make([]float64, xr)
	dr, _ := af.xMem.Dims()
	dCol := make([]float64, dr)
	// create input matrix and target vector
	// shift column
	for i := af.order - 1; i > 0; i-- {
		mat.Col(xCol, i-1, af.xMem)
		af.xMem.SetCol(i, xCol)
		mat.Col(dCol, i-1, af.dMem)
		af.dMem.SetCol(i, dCol)
	}
	// estimate output and error
	wd := mat.NewDense(1, len(af.w), af.w)
	af.yMem.Mul(wd, af.xMem.T())
	af.eMem.Sub(af.dMem, af.yMem)

	// update
	dw1 := mat.NewDense(af.order, af.order, nil)
	dw1.Mul(af.xMem.T(), af.xMem)
	dw1.Add(dw1, af.epsIDE)
	dw2 := mat.NewDense(af.order, af.order, nil)
	err := dw2.Solve(dw1, af.ide)
	if err != nil {
		return  err
	}
	dw3 := mat.NewDense(1, af.order, nil)
	dw3.Mul(af.eMem, dw2)
	dw := mat.NewDense(1, af.n, nil)
	dw.Scale(af.mu, dw)
	floats.Add(af.w, dw.RawRowView(0))
	return nil
}

func (af *FiltAP) Run(d []float64, x [][]float64) ([]float64, []float64, [][]float64, error) {
	//measure the data and check if the dimension agree
	N := len(x)
	if len(d) != N {
		return nil, nil, nil, errors.New("the length of slice d and x must agree.")
	}
	af.n = len(x[0])
	//af.wHistory = make([][]float64, N)

	y := make([]float64, N)
	e := make([]float64, N)

	xr, _ := af.xMem.Dims()
	xCol := make([]float64, xr)
	dr, _ := af.xMem.Dims()
	dCol := make([]float64, dr)
	//adaptation loop
	for i := 0; i < N; i++ {
		//af.wHistory[i] = af.w
		af.wHistory.SetRow(i, af.w)

		// create input matrix and target vector
		// shift column
		for i := af.order - 1; i > 0; i-- {
			mat.Col(xCol, i-1, af.xMem)
			af.xMem.SetCol(i, xCol)
			mat.Col(dCol, i-1, af.dMem)
			af.dMem.SetCol(i, dCol)
		}

		// estimate output and error
		wd := mat.NewDense(1, len(af.w), af.w)
		af.yMem.Mul(wd, af.xMem.T())
		af.eMem.Sub(af.dMem, af.yMem)
		y[i] = af.yMem.At(0, 0)
		e[i] = af.eMem.At(0, 0)

		// update
		dw1 := mat.NewDense(af.order, af.order, nil)
		dw1.Mul(af.xMem.T(), af.xMem)
		dw1.Add(dw1, af.epsIDE)
		dw2 := mat.NewDense(af.order, af.order, nil)
		err := dw2.Solve(dw1, af.ide)
		if err != nil {
			return nil, nil, nil, err
		}
		dw3 := mat.NewDense(1, af.order, nil)
		dw3.Mul(af.eMem, dw2)
		dw := mat.NewDense(1, af.n, nil)
		dw.Scale(af.mu, dw)
		floats.Add(af.w, dw.RawRowView(0))
	}
	wHistory := make([][]float64, af.n)
	for i := 0; i < af.n; i++ {
		wHistory[i] = af.wHistory.RawRowView(i)
	}
	return y, e, wHistory, nil
}

