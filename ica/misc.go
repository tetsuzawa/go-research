package ica

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func NewRandSlice(n int) []float64 {
	rs := make([]float64, n)
	for i:=0;i<n;i++{
		rs[i] = rand.Float64()
	}
	return rs
}

func NewRandVector(n int) *mat.Dense {
	rs := make([]float64, n)
	for i:=0;i<n;i++{
		rs[i] = rand.Float64()
	}
	vec := mat.NewDense(1, n, rs)
	return vec
}


func NormalizeMat(x *mat.Dense) *mat.Dense{
	if mat.Sum(x) < 0 {
		x.Scale(-1, x)
	}
	x.Scale(mat.Norm(x, 1), x)
	return x
}
