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
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type ADFInterface interface {
	InitWeights(w interface{}, n int) error
	Predict(x []float64) (y float64)
	//PreTrainedRun(d []float64, x [][]float64, nTrain float64, epochs int) (y, e []float64, w [][]float64, err error)
	Adapt(d float64, x []float64)
	Run(d []float64, x [][]float64) ([]float64, []float64, [][]float64, error)
	//ExploreLearning(d []float64, x [][]float64, muStart, muEnd float64, steps int,
	//	nTrain float64, epochs int, criteria string, targetW []float64) ([]float64, []float64, error)
	CheckFloatParam(p, low, high float64, name string) (float64, error)
	CheckIntParam(p, low, high int, name string) (int, error)
	SetMu(mu float64)
	GetParams() (int, float64, []float64)
}

//AdaptiveFilter is base struct for adaptive filter structs
//It puts together some functions used by all adaptive filters.
type AdaptiveFilter struct {
	w  *mat.Dense
	n  int
	mu float64
}

func newAdaptiveFilter(n int, mu float64, w interface{}) (ADFInterface, error) {
	var err error
	p := new(AdaptiveFilter)
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

func Must(adf ADFInterface, err error) ADFInterface {
	if err != nil {
		panic(err)
	}
	return adf
}

//PreTrainedRun sacrifices part of the data for few epochs of learning.
//`d`: desired value
//`x`: input matrix (samples x input arrays). rows are samples and  columns are features.
//`nTrain`: train to test ratio (float), default value is 0.5
//          (that means 50% of data is used for training)
//`epochs`: number of training epochs (int), default value is 1.
//          This number describes how many times the training will be repeated
//          on dedicated part of data.
func PreTrainedRun(af ADFInterface, d []float64, x [][]float64, nTrain float64, epochs int) (y, e []float64, w [][]float64, err error) {
	var nTrainI = int(float64(len(d)) * nTrain)
	//train
	for i := 0; i < epochs; i++ {
		_, _, _, err = af.Run(d[:nTrainI], x[:][:nTrainI])
		if err != nil {
			return nil, nil, nil, err
		}
	}
	//run
	y, e, w, err = af.Run(d[:nTrainI], x[:nTrainI])
	if err != nil {
		return nil, nil, nil, err
	}
	return y, e, w, nil
}

//ExploreLearning tests what learning rate is the best.
//
//* `d` : desired value.
//* `x` : input matrix.
//* `muStart` : starting learning rate.
//* `muEnd` : final learning rate.
//* `steps` : how many learning rates should be tested between `muStart`
//			  and `muEnd`.
//* `nTrain` : train to test ratio , default value is 0.5.
//			   (that means 50% of data is used for training)
//* `epochs` : number of training epochs , default value is 1.
//			   This number describes how many times the training will be repeated
//			   on dedicated part of data.
//* `criteria` : how should be measured the mean error,
//				 default value is "MSE".
//* `target_w` : target weights, default value is False.
//				 If False, the mean error is estimated from prediction error.
//				 If an array is provided, the error between weights and `target_w`
//				 is used.
func ExploreLearning(af ADFInterface, d []float64, x [][]float64, muStart, muEnd float64, steps int,
	nTrain float64, epochs int, criteria string, targetW []float64) ([]float64, []float64, error) {
	mus := LinSpace(muStart, muEnd, steps)
	es := make([]float64, len(mus))
	zeros := make([]float64, int(float64(len(x))*nTrain))
	for i, mu := range mus {
		//init
		err := af.InitWeights("zeros", len(x[0]))
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to init weights at InitWights()")
		}
		af.SetMu(mu)
		//run
		_, e, _, err := PreTrainedRun(af, d, x, nTrain, epochs)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to pre train at PreTrainedRun()")
		}
		es[i], err = GetMeanError(e, zeros, criteria)
		//fmt.Println(es[i])
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to get mean error at GetMeanError()")
		}
	}
	return es, mus, nil
}

func (af *AdaptiveFilter) GetParams() (int, float64, []float64) {
	return af.n, af.mu, af.w.RawRowView(0)
}

//InitWeights initialises the adaptive weights of the filter.
//
//`w`: initial weights of filter. Possible values are
//* "random": create random weights
//* "zeros": create zero value weights
//
//`n`: size of filter (int) - number of filter coefficients.
func (af *AdaptiveFilter) InitWeights(w interface{}, n int) error {
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
func (af *AdaptiveFilter) Predict(x []float64) (y float64) {
	y = floats.Dot(af.w.RawRowView(0), x)
	return y
}

//Override to use this func.
func (af *AdaptiveFilter) Adapt(d float64, x []float64) {
	//TODO
}

//Override to use this func.
func (af *AdaptiveFilter) Run(d []float64, x [][]float64) ([]float64, []float64, [][]float64, error) {
	//TODO
	//measure the data and check if the dimension agree
	N := len(x)
	if len(d) != N {
		return nil, nil, nil, errors.New("the length of slice d and x must agree")
	}
	af.n = len(x[0])

	y := make([]float64, N)
	e := make([]float64, N)
	w := make([]float64, af.n)
	ws := make([][]float64, N)
	//adaptation loop
	for i := 0; i < N; i++ {
		w = af.w.RawRowView(0)
		y[i] = floats.Dot(w, x[i])
		e[i] = d[i] - y[i]
		ws[i] = w
	}
	return y, e, ws, nil
}

//CheckFloatParam check if the value of the given parameter
//is in the given range and a float.
func (af *AdaptiveFilter) CheckFloatParam(p, low, high float64, name string) (float64, error) {
	if low <= p && p <= high {
		return p, nil
	} else {
		err := fmt.Errorf("parameter %v is not in range <%v, %v>", name, low, high)
		return 0, err
	}
}

//CheckIntParam check if the value of the given parameter
//is in the given range and a int.
func (af *AdaptiveFilter) CheckIntParam(p, low, high int, name string) (int, error) {
	if low <= p && p <= high {
		return p, nil
	} else {
		err := fmt.Errorf("parameter %v is not in range <%v, %v>", name, low, high)
		return 0, err
	}
}

//SetMu set a update param mu.
func (af *AdaptiveFilter) SetMu(mu float64) {
	af.mu = mu
}
