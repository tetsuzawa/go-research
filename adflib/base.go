package adflib

import (
	"errors"
	"math/rand"
	"reflect"
	"time"
)

type AdaptiveFilter struct {
	w []float64
	n int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewRandn returns random value. stddev 0.5, mean 0.
func NewRandn() float64 {
	return rand.NormFloat64()*0.5 + 0
}

func (af *AdaptiveFilter) InitWeghts(w interface{}, n int) error {
	switch v := w.(type) {
	case string:
		if v == "random" {
			w := make([]float64, n)
			for i := 0; i < n; i++ {
				w[i] = NewRandn()
			}
			af.w = w
			return nil
		} else if v == "zeros" {
			w := make([]float64, n)
			af.w = w
			return nil
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
}
