package ica

import (
	"errors"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
	"math"
)

type ICA struct {
	xMat    *mat.Dense
	sampNum int
	sigNum  int
}

func NewICA(x [][]float64) *ICA {
	ica := new(ICA)
	x1d := make([]float64, len(x)*len(x[0]))
	for i, inSl := range x {
		for j, v := range inSl {
			x1d[len(x)*i+j] = v
		}
	}
	ica.xMat = mat.NewDense(len(x), len(x[0]), x1d)
	ica.sampNum, ica.sigNum = ica.xMat.Dims()
	return ica
}

func (ica *ICA) CalcICA() ([][]float64, error) {
	ica.fit()
	z, err := ica.whiten()
	if err != nil {
		return nil, err
	}
	y := ica.analyze(z)
	return y, nil
}

func (ica *ICA) fit() {
	sig := make([]float64, ica.sampNum)
	for i := 0; i < ica.sigNum; i++ {
		mat.Col(sig, i, ica.xMat)
		sigMean := floats.Sum(sig) / float64(len(sig))
		for j := 0; j < ica.sampNum; j++ {
			sig[j] -= sigMean
		}
		ica.xMat.SetCol(i, sig)
	}
}

func (ica *ICA) whiten() (*mat.Dense, error) {
	sigma := mat.NewSymDense(ica.sigNum, nil)
	stat.CorrelationMatrix(sigma, ica.xMat, nil)
	var eigsym mat.EigenSym
	ok := eigsym.Factorize(sigma, true)
	if !ok {
		return nil, errors.New("symmetric eigendecomposition failed")
	}

	// eigenvalues of sigma
	D := eigsym.Values(nil)

	// eigenvectors of sigma
	var V *mat.Dense
	eigsym.VectorsTo(V)

	DhSl := make([]float64, ica.sigNum*ica.sigNum)
	for i := 0; i < ica.sigNum; i++ {
		DhSl[i*(ica.sigNum+1)] = math.Pow(D[i], -1/2)
	}
	Dh := mat.NewDense(ica.sigNum, ica.sigNum, DhSl)
	retMat := produ



}
