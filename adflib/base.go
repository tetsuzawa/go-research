/*
This package is designed to simplify adaptive signal processing tasks
within golang (filtering, prediction, reconstruction, classification).
For code optimisation, this library uses gonum/floats for array operations.

This package is created with reference to https://github.com/matousc89/padasip.
*/
package adflib

import (
	"errors"
	"fmt"
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
	PreTrainedRun(d []float64, x [][]float64, nTrain float64, epochs int) (y, e []float64, w [][]float64, err error)
	Adapt(d float64, x []float64)
	Run(d []float64, x [][]float64) ([]float64, []float64, [][]float64, error)
	ExploreLearning(d []float64, x [][]float64, muStart, muEnd float64, steps int,
		nTrain float64, epochs int, criteria string, targetW []float64) ([]float64, error)
	CheckFloatParam(p, low, high float64, name string) (float64, error)
	CheckIntParam(p, low, high int, name string) (int, error)
}

//AdaptiveFilter is base struct for adaptive filter structs
//It puts together some functions used by all adaptive filters.
type AdaptiveFilter struct {
	w  *mat.Dense
	n  int
	mu float64
}

// NewRandn returns random value. stddev 0.5, mean 0.
func NewRandn(stddev, mean float64) float64 {
	return rand.NormFloat64()*stddev + mean
}

func LinSpace(start, end float64, n int) []float64 {
	res := make([]float64, n)
	if n == 1 {
		res[0] = end
		return res
	}
	delta := (end - start) / (float64(n) - 1)
	for i := 0; i < n; i++ {
		res[i] = start + (delta * float64(i))
	}
	return res
}

func newAdaptiveFilter(w [][]float64, n int, mu float64) ADFInterface {
	wMat := mat.NewDense(len(w), len(w[0]), Floor(w))
	return &AdaptiveFilter{
		w:  wMat,
		n:  n,
		mu: mu,
	}
}

func Must(adf ADFInterface, err error) ADFInterface {
	if err != nil {
		panic(err)
	}
	return adf
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

//PreTrainedRun sacrifices part of the data for few epochs of learning.
//`d`: desired value
//`x`: input matrix (samples x input arrays). rows are samples and  columns are features.
//`nTrain`: train to test ratio (float), default value is 0.5
//          (that means 50% of data is used for training)
//`epochs`: number of training epochs (int), default value is 1.
//          This number describes how many times the training will be repeated
//          on dedicated part of data.
func (af *AdaptiveFilter) PreTrainedRun(d []float64, x [][]float64, nTrain float64, epochs int) (y, e []float64, w [][]float64, err error) {
	var nTrainI = int(float64(len(d)) * nTrain)
	for i := 0; i < epochs; i++ {
		_, _, _, err = af.Run(d[:nTrainI], x[:][:nTrainI])
		if err != nil {
			return nil, nil, nil, err
		}
	}
	y, e, w, err = af.Run(d[:nTrainI], x[:nTrainI])
	if err != nil {
		return nil, nil, nil, err
	}
	return y, e, w, nil
}

//Override to use this func.
func (af *AdaptiveFilter) Adapt(d float64, x []float64) {
	//TODO
}

//Override to use this func.
func (af *AdaptiveFilter) Run(d []float64, x [][]float64) ([]float64, []float64, [][]float64, error) {
	//TODO
	return nil, nil, nil, nil
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
func (af *AdaptiveFilter) ExploreLearning(d []float64, x [][]float64, muStart, muEnd float64, steps int,
	nTrain float64, epochs int, criteria string, targetW []float64) ([]float64, error) {
	mus := LinSpace(muStart, muEnd, steps)
	es := make([]float64, len(mus))
	zeros := make([]float64, len(mus))
	for i, mu := range mus {
		//init
		err := af.InitWeights("zeros", 0)
		if err != nil {
			return nil, err
		}
		af.mu = mu
		//run
		_, e, _, err := af.PreTrainedRun(d, x, nTrain, epochs)
		if err != nil {
			return nil, err
		}
		es[i], err = GetMeanError(e, zeros, criteria)
		if err != nil {
			return nil, err
		}
	}
	return es, nil
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
