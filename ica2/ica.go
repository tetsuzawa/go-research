package ica

import (
	"errors"
	"github.com/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
	"math"
)

func g(x float64) float64 {
	return math.Tanh(x)
}

func gDer(x float64) float64 {
	return 1 - g(x)*g(x)
}

func center(x *mat.Dense) *mat.Dense {
	r, c := x.Dims()
	var xs = make([]float64, c)
	var xMeans = make([]float64, r)
	for i := 0; i < r; i++ {
		xs = x.RawRowView(i)
		xMeans[i] = floats.Sum(xs) / float64(c)
	}
	return mat.NewDense(r, 1, xMeans)
}

func whitening(x *mat.Dense) (*mat.Dense, error) {
	r, c := x.Dims()
	cov := mat.NewSymDense(r, nil)
	stat.CorrelationMatrix(cov, x, nil)
	var eigsym mat.EigenSym
	ok := eigsym.Factorize(cov, true)
	if !ok {
		return nil, errors.New("symmetric eigendecomposition failed")
	}
	// eigenvalues of cov
	d := eigsym.Values(nil)
	// eigenvectors of cov
	E := eigsym.VectorsTo(nil)
	D := NewDiagMat(d, r)
	var DInv = mat.NewDense(r, r, nil)
	err := DInv.Inverse(D)
	if err != nil {
		return nil, err
	}
	for i := 0; i < r; i++ {
		DInv.Set(i, i, math.Sqrt(DInv.At(i, i)))
	}
	var XWhiten = mat.NewDense(r, c, nil)
	XWhiten.Product(E, DInv, E.T(), x)
	return XWhiten, err
}
