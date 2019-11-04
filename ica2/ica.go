package ica

import (
	"errors"
	"fmt"
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
	for i := 0; i < r; i++ {
		xs = X.RawRowView(i)
		floats.AddConst(-(floats.Sum(xs) / float64(c)), xs)
	}
	return X
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
	var auxSl3 = make([]float64, 1*c)
	aux1 := mat.NewDense(1, c, nil)
	aux1.Mul(w, X)
	row1 := aux1.RawRowView(0)
	for i := 0; i < c; i++ {
		auxSl1[i] = g(row1[i])
		auxSl2[i] = gDer(row1[i])
	}

	aux3 := mat.NewDense(r, c, nil)
	// diagonal matrix
	//aux2 := NewDiagMat(auxSl1, c)
	for i := 0; i < r; i++ {
		copy(auxSl3, X.RawRowView(i))
		floats.Mul(auxSl3, auxSl1)
		aux3.SetRow(i, auxSl3)
	}
	//aux3.Mul(X, aux2)
	// [r, 1]
	aux4 := ColMeanVector(aux3)
	// diagonal matrix
	//aux5 := mat.NewDense(1, c, auxSl2)
	//aux5 := NewDiagMat(auxSl2, c)
	//aux6 is scalar
	aux6 := SliceMean(auxSl2)
	aux7 := mat.NewDense(1, r, nil)
	aux7.Scale(aux6, w)
	wNew := mat.NewDense(1, r, nil)
	wNew.Sub(aux4.T(), aux7)
	aux8 := make([]float64, 3)
	copy(aux8, wNew.RawRowView(0))
	for i, v := range aux8 {
		aux8[i] = v * v
	}
	wNew.Scale(1/math.Sqrt(floats.Sum(aux8)), wNew)
	return wNew
}

func ICA(X *mat.Dense, iter int, tolerance float64) (*mat.Dense, error) {
	r, c := X.Dims()
	X = center(X)
	X, err := Whitening(X)
	if err != nil {
		return nil, err
	}
	componentsNR, _ := X.Dims()
	W := mat.NewDense(componentsNR, componentsNR, nil)

	var w = mat.NewDense(1, componentsNR, nil)
	var wNew = mat.NewDense(1, componentsNR, nil)
	// aux
	var aMat1 = mat.NewDense(1, componentsNR, nil)
	var aMat2 = mat.NewDense(1, componentsNR, nil)
	var distance float64
	for i := 0; i < componentsNR; i++ {
		w = NewRandVector(componentsNR)

		for j := 0; j < iter; j++ {
			wNew = CalcNewW(w, X)
			fmt.Printf("Calculating... %d%%\r", (i*iter+j+1)*100/(componentsNR*iter))
			if i >= 1 {
				//TODO
				Wi := W.Slice(0, i, 0, componentsNR)
				//wnWT := mat.NewDense(1, i, nil)
				//wnWT.Mul(wNew, Wi.T())
				//aMat1.Mul(wnWT, Wi)
				//wNew.Sub(wNew, aMat1)
				aMat1.Product(wNew, Wi.T(), Wi)
				wNew.Sub(wNew, aMat1)
			}
			aMat2.MulElem(w, wNew)
			distance = math.Abs(math.Abs(ElemSum(aMat2) - 1))
			w = wNew

			if distance < tolerance {
				break
			}
		}
		W.SetRow(i, w.RawRowView(0))
	}

	S := mat.NewDense(r, c, nil)
	S.Mul(W, X)
	fmt.Println("complete")
	return S, nil
}
