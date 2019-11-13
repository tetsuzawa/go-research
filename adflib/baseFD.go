/*
This package is designed to simplify adaptive signal processing tasks
within golang (filtering, prediction, reconstruction, classification).
For code optimisation, this library uses gonum/floats for array operations.

This package is created with reference to https://github.com/matousc89/padasip.
*/
package adflib

import (
	"fmt"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/mat"
)

type FDADFInterface interface {
	InitWeights(w interface{}, n int) error
	Predict(x []float64) (y []float64)
	Adapt(d []float64, x []float64)
	Run(d [][]float64, x [][]float64) ([][]float64, [][]float64, [][]float64, error)
	CheckFloatParam(p, low, high float64, name string) (float64, error)
	CheckIntParam(p, low, high int, name string) (int, error)
	SetMu(mu float64)
	GetParams() (int, float64, []float64)
}

//FDAdaptiveFilter is base struct for frequency domain adaptive filter structs
//It puts together some functions used by all adaptive filters.
type FDAdaptiveFilter struct {
	w  *mat.Dense
	n  int
	mu float64
}

func newFDAdaptiveFilter(n int, mu float64, w interface{}) (FDADFInterface, error) {
	var err error
	p := new(FDAdaptiveFilter)
	p.n = n
	p.mu, err = p.CheckFloatParam(mu, 0, 1000, "mu")
	if err != nil {
		return nil, err
	}
	err = p.InitWeights(w, n)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (af *FDAdaptiveFilter) GetParams() (int, float64, []float64) {
	return af.n, af.mu, af.w.RawRowView(0)
}

//InitWeights initialises the adaptive weights of the filter.
//
//`w`: initial weights of filter. Possible values are
//* "random": create random weights
//* "zeros": create zero value weights
//
//`n`: size of filter (int) - number of filter coefficients.
func (af *FDAdaptiveFilter) InitWeights(w interface{}, n int) error {
	if n <= 0 {
		n = af.n
	}
	switch v := w.(type) {
	case string:
		if v == "random" {
			w := make([]float64, n)
			for i := 0; i < n; i++ {
				w[i] = NewRandn(0.5, 0)
			}
			af.w = mat.NewDense(1, n, w)
		} else if v == "zeros" {
			w := make([]float64, n)
			af.w = mat.NewDense(1, n, w)
		} else {
			return errors.New("impossible to understand the w")
		}
	case []float64:
		if len(v) != n {
			return errors.New("length of w is different from n")
		}
		af.w = mat.NewDense(1, n, v)
	default:
		return errors.New(`args w must be "random" or "zeros" or []float64{...}`)
	}
	return nil
}

//Predict calculates the new output value `y` from input array `x`.
func (af *FDAdaptiveFilter) Predict(x []float64) (y []float64) {
	//TODO
	//y = floats.Dot(af.w.RawRowView(0), x)
	//return y
	copy(y, x)
	return
}

//Override to use this func.
func (af *FDAdaptiveFilter) Adapt(d []float64, x []float64) {
	//TODO
}

//Override to use this func.
func (af *FDAdaptiveFilter) Run(d [][]float64, x [][]float64) ([][]float64, [][]float64, [][]float64, error) {
	//TODO
	return nil, nil, nil, nil
}

//CheckFloatParam check if the value of the given parameter
//is in the given range and a float.
func (af *FDAdaptiveFilter) CheckFloatParam(p, low, high float64, name string) (float64, error) {
	if low <= p && p <= high {
		return p, nil
	} else {
		err := fmt.Errorf("parameter %v is not in range <%v, %v>", name, low, high)
		return 0, err
	}
}

//CheckIntParam check if the value of the given parameter
//is in the given range and a int.
func (af *FDAdaptiveFilter) CheckIntParam(p, low, high int, name string) (int, error) {
	if low <= p && p <= high {
		return p, nil
	} else {
		err := fmt.Errorf("parameter %v is not in range <%v, %v>", name, low, high)
		return 0, err
	}
}

//SetMu set a update param mu.
func (af *FDAdaptiveFilter) SetMu(mu float64) {
	af.mu = mu
}
