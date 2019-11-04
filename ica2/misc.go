package ica

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
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

// ColMeanVector returns a means of matrix row
//	[ ],
//	[ ],
//	[ ]
func ColMeanVector(X *mat.Dense) *mat.Dense {
	r, c := X.Dims()
	var xs = make([]float64, c)
	var xMeans = make([]float64, r)
	for i := 0; i < r; i++ {
		xs = X.RawRowView(i)
		xMeans[i] = floats.Sum(xs) / float64(c)
	}
	return mat.NewDense(r, 1, xMeans)
}

// RowMeanVector returns a means of matrix col
// 	[ , , ,]
func RowMeanVector(X *mat.Dense) *mat.Dense {
	r, c := X.Dims()
	var xs = make([]float64, r)
	var xMeans = make([]float64, c)
	for i := 0; i < c; i++ {
		mat.Col(xs, i, X)
		xMeans[i] = floats.Sum(xs) / float64(r)
	}
	return mat.NewDense(1, c, xMeans)
}

func SliceMean(fs []float64) float64 {
	return floats.Sum(fs) / float64(len(fs))
}
