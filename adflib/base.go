package adflib

import (
	"errors"
	"fmt"
	"github.com/tetsuzawa/go-research/adflib/misc"
	"gonum.org/v1/gonum/floats"
	"math/rand"
	"time"
)

type ADFInterface interface {
	InitWeights() error
	Predict() float64
	PreTrainedRun() ([]float64, []float64, []float64)
	Run() ([]float64, []float64, []float64)
	ExploreLearning() ([]float64, error)
	CheckFloatParam() (float64, error)
	CheckIntParam() (int, error)
}

//AdaptiveFilter is base struct for adaptive filter structs.
//It puts together some functions used by all adaptive filters.
type AdaptiveFilter struct {
	w  []float64
	n  int
	mu float64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewRandn returns random value. stddev 0.5, mean 0.
func NewRandn() float64 {
	return rand.NormFloat64()*0.5 + 0
}

func linspace(start, end float64, n int) ([]float64) {
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
				w[i] = NewRandn()
			}
			af.w = w
		} else if v == "zeros" {
			w := make([]float64, n)
			af.w = w
		} else {
			return errors.New("Impossible to understand the w")
		}
	case []float64:
		if len(v) != n {
			return errors.New("length of w is different from n")
		}
		af.w = v
	default:
		return errors.New(`args w must be "random" or "zeros" or []float64{...}`)
	}
	return nil
}

//Predict calculates the new output value `y` from input array `x`.
func (af *AdaptiveFilter) Predict(x []float64) float64 {
	var y float64
	y = floats.Dot(af.w, x)
	return y
}

//PreTrainedRun sacrifices part of the data for few epochs of learning.
//`d`: desired value
//`x`: input matrix (samples x input arrays)
//`nTrain`: train to test ratio (float), default value is 0.5
//          (that means 50% of data is used for training)
//`epochs`: number of training epochs (int), default value is 1.
//          This number describes how many times the training will be repeated
//          on dedicated part of data.
func (af *AdaptiveFilter) PreTrainedRun(d, x []float64, nTrain float64, epochs int) (y, e, w []float64) {
	var Ntrain = int(float64(len(d)) * nTrain)
	for i := 0; i < epochs; i++ {
		af.Run(d[:Ntrain], x[:Ntrain])
	}
	y, e, w = af.Run(d[:Ntrain], x[:Ntrain])
	return y, e, w
}

//Override to use this func.
func (af *AdaptiveFilter) Run(d, x []float64) (y, e, w []float64) {
	//TODO
	return nil, nil, nil
}

func (af *AdaptiveFilter) ExploreLearning(d, x []float64, muStart, muEnd float64, steps int,
	nTrain float64, epochs int, criteria string, targetW []float64) ([]float64, error) {
	mus := linspace(muStart, muEnd, steps)
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
		_, e, _ := af.PreTrainedRun(d, x, nTrain, epochs)
		es[i], err = misc.GetMeanError(e, zeros, criteria)
		if err != nil {
			return nil, err
		}
	}
	return es, nil
}

func (af *AdaptiveFilter) CheckFloatParam(p, low, high float64, name string) (float64, error) {
	if low <= p && p <= high {
		return p, nil
	} else {
		err := fmt.Errorf("Parameter %v is not in range <%v, %v>")
		return 0, err
	}
}
func (af *AdaptiveFilter) CheckIntParam(p, low, high int, name string) (int, error) {
	if low <= p && p <= high {
		return p, nil
	} else {
		err := fmt.Errorf("Parameter %v is not in range <%v, %v>")
		return 0, err
	}
}
