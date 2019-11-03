package ica

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewRandSlice(n int) []float64 {
	rs := make([]float64, n)
	for i := 0; i < n; i++ {
		rs[i] = rand.Float64()
	}
	return rs
}

func NewRandVector(n int) *mat.Dense {
	rs := make([]float64, n)
	for i := 0; i < n; i++ {
		rs[i] = rand.Float64()
	}
	vec := mat.NewDense(1, n, rs)
	return vec
}

func NewDiagMat(fs []float64, n int) *mat.Dense {
	if fs == nil {
		fs := make([]float64, n)
		for i, _ := range fs {
			fs[i] = 1
		}
	}
	dSl := make([]float64, n*n)
	for i := 0; i < n; i++ {
		dSl[i*(n+1)] = fs[i]
	}
	return mat.NewDense(n, n, dSl)
}

func NormalizeMat(x *mat.Dense) *mat.Dense {
	if mat.Sum(x) < 0 {
		x.Scale(-1, x)
	}
	x.Scale(mat.Norm(x, 1), x)
	return x
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
