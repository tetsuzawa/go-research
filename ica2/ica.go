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

func center(X *mat.Dense) *mat.Dense {
	r, c := X.Dims()
	var xs = make([]float64, c)
	var xMeans = make([]float64, r)
	for i := 0; i < r; i++ {
		xs = X.RawRowView(i)
		xMeans[i] = floats.Sum(xs) / float64(c)
	}
	return mat.NewDense(r, 1, xMeans)
}

func Whitening(X *mat.Dense) (*mat.Dense, error) {
	r, c := X.Dims()
	cov := mat.NewSymDense(r, nil)
	stat.CovarianceMatrix(cov, X.T(), nil)
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
	XWhiten.Product(E, DInv, E.T(), X)
	return XWhiten, err
}

//CalcNewW calculate new w.
//w is *mat.Dense [1, r]
//X is *mat.Dense [r, c]: [signul num, sample num]
func CalcNewW(w, X *mat.Dense) *mat.Dense {
	r, c := X.Dims()
	var auxSl1 = make([]float64, 1*c)
	var auxSl2 = make([]float64, 1*c)
	aux1 := mat.NewDense(1, c, nil)
	aux1.Mul(w, X)
	row1 := aux1.RawRowView(0)
	for i := 0; i < c; i++ {
		auxSl1[i] = g(row1[i])
		auxSl2[i] = gDer(row1[i])
	}
	// diagonal matrix
	//aux2 := mat.NewDense(1, c, auxSl1)
	aux2 := NewDiagMat(auxSl1, c)
	aux3 := mat.NewDense(r, c, nil)
	aux3.Mul(X, aux2)
	// [r, 1]
	aux4 := ColMeanVector(aux3)
	// diagonal matrix
	//aux5 := mat.NewDense(1, c, auxSl2)
	//aux5 := NewDiagMat(auxSl2, c)
	//aux6 is scalar
	aux6 := SliceMean(auxSl2)
	//aux6 := RowMeanVector(aux5)
	aux7 := mat.NewDense(1, r, nil)
	aux7.Scale(aux6, w)
	wNew := mat.NewDense(1, r, nil)
	wNew.Sub(aux4.T(), aux7)
	aux8 := make([]float64, 3)
	copy(aux8, wNew.RawRowView(0))
	for i, v := range aux8{
		aux8[i] = v * v
	}
	//aux8.Pow(wNew, 2)
	//aux8 := mat.NewDense(1, r, row2)
	wNew.Scale(1/math.Sqrt(floats.Sum(aux8)), wNew)
	return wNew

	/* 行列と間違えた
	var auxSl1 = make([]float64, r*c)
	var auxSl2 = make([]float64, r*c)
	aux1 := mat.NewDense(1, c, nil)
	aux1.Mul(w.T(), X)
	for i := 0; i < r; i++ {
		row1 := aux1.RawRowView(i)
		for j := 0; j < c; j++ {
			auxSl1[c*i+j] = g(row1[j])
			auxSl2[c*i+j] = gDer(row1[j])
		}
	}
	aux2 := mat.NewDense(r, c, auxSl1)
	aux3 := mat.NewDense(r, c, nil)
	aux3.Mul(X, aux2)
	aux4 := ColMeanVector(aux3)
	aux5 := mat.NewDense(r, c, auxSl2)
	aux6 := RowMeanVector(aux5)
	aux7 :=
		fmt.Println(aux4)
	//wNew :=
	*/
}
