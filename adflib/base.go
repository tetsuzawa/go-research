package adflib

import (
	"errors"
	"gonum.org/v1/gonum/floats"
	"math/rand"
	"time"
)

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

func (af *AdaptiveFilter) InitWeghts(w interface{}, n int) error {
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

func (af *AdaptiveFilter) Predict(x []float64) float64 {
	var y float64
	y = floats.Dot(af.w, x)
	return y
}

func (af *AdaptiveFilter) PreTrainedRun(d, x []float64, ntrain float64, epochs int) (y, e float64, w []float64) {
	var Ntrain = int(float64(len(d)) * ntrain)
	for i := 0; i < epochs; i++ {
		af.Run(d[:Ntrain], x[:Ntrain])
	}
	y, e, w = af.Run(d[:Ntrain], x[:Ntrain])
	return y, e, w
}

//Override to use this func.
func (af *AdaptiveFilter) Run(d, x []float64) (y, e float64, w []float64) {
	//TODO
	return 0, 0, nil
}

func (af *AdaptiveFilter) ExploreLearning(d, x []float64, muStart, muEnd float64, steps int,
	nTrain float64, epochs int, criteria string, targetW bool) error {
	mus := linspace(muStart, muEnd, steps)
	for i, mu := range mus {
		//init
		err := af.InitWeghts("zeros", 0)
		if err != nil {
			return err
		}
		af.mu = mu
		//run
		y, e, w := af.PreTrainedRun(d, x, nTrain, epochs)

	}
}


