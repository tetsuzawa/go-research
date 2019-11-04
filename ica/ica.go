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
			x1d[len(x)*j+i] = v
		}
	}
	ica.xMat = mat.NewDense(len(x[0]), len(x), x1d)
	ica.sampNum, ica.sigNum = ica.xMat.Dims()
	return ica
}

func (ica *ICA) CalcICA() ([][]float64, error) {
	ica.fit()
	z, err := ica.whiten()
	if err != nil {
		return nil, err
	}
	yMat := ica.analyze(z)
	y := make([][]float64, ica.sigNum)
	for i := 0; i < ica.sigNum; i++ {
		mat.Col(y[i], i, yMat)
	}
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
	V = eigsym.VectorsTo(nil)

	DhSl := make([]float64, ica.sigNum*ica.sigNum)
	for i := 0; i < ica.sigNum; i++ {
		//TODO tmp val
		tmp := math.Pow(D[i], -1./2.)
		DhSl[i*(ica.sigNum+1)] = tmp
	}
	Dh := mat.NewDense(ica.sigNum, ica.sigNum, DhSl)
	z := mat.NewDense(ica.sigNum, ica.sampNum, nil)
	V.Product(V, Dh, V.T())
	z.Mul(V, ica.xMat.T())
	//zt := z.T()
	//zt := mat.NewDense(ica.sampNum, ica.sigNum, z.T())

	return z, nil
}

func (ica *ICA) normalize() {
	if mat.Sum(ica.xMat) < 0 {
		ica.xMat.Scale(-1, ica.xMat)
	}
	ica.xMat.Scale(mat.Norm(ica.xMat, 1), ica.xMat)
}
func (ica *ICA) analyze(z *mat.Dense) *mat.Dense {
	c := ica.sigNum
	//TODO
	//W := make([]float64, c)

	// aux matrix
	aMat := mat.NewDense(ica.sigNum, ica.sigNum, nil)
	// execute analysis for count of observations
	for i := 0; i < c; i++ {
		//wSl := NewRandSlice(c)
		wVec := NewRandVector(c) //(1, 3)
		//wSl = floats.Norm(wSl, )
		wVec = NormalizeMat(wVec)
		//wDia := NewDiagMat(wVec.RawRowView(0), c)
		//means := make([][]float64, c)

		for {
			aMat.Copy(z)
			aMat.Mul(z, wVec.T())
			aMat.MulElem(aMat, aMat)
			aMat.MulElem(aMat, aMat)
			break

			//TODO
			//for j := 0; j < c; j++ {
			//	means[j] = cVec.
			//}
			//
			//wVec.T
			//wVecPre := wVec
			//wVec =
		}
	}
	y := mat.NewDense(ica.sampNum, ica.sigNum, nil)
	//y.Mul(W,z)
	return y
}
