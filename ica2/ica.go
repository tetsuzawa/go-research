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
	var xs = make([]float64, r)
	for i := 0; i < c; i++ {
		//xs = X.RawRowView(i)
		mat.Col(xs, i, X)
		floats.AddConst(-(floats.Sum(xs) / float64(r)), xs)
		X.SetCol(i, xs)
	}
	return X
	//var xs = make([]float64, c)
	//for i := 0; i < r; i++ {
	//	xs = X.RawRowView(i)
	//	floats.AddConst(-(floats.Sum(xs) / float64(c)), xs)
	//}
	//return X
}

func Whiten(X *mat.Dense) (*mat.Dense, error) {
	r, c := X.Dims()
	cov := mat.NewSymDense(c, nil)
	stat.CovarianceMatrix(cov, X, nil)
	var eigsym mat.EigenSym
	ok := eigsym.Factorize(cov, true)
	if !ok {
		return nil, errors.New("symmetric eigendecomposition failed")
	}
	// eigenvalues of cov
	d := eigsym.Values(nil)
	// eigenvectors of cov
	E := eigsym.VectorsTo(nil)
	D := NewDiagMat(d, c)
	var DInv = mat.NewDense(c, c, nil)
	err := DInv.Inverse(D)
	if err != nil {
		return nil, err
	}
	for i := 0; i < c; i++ {
		DInv.Set(i, i, math.Sqrt(DInv.At(i, i)))
	}
	var XWhitenT = mat.NewDense(c, r, nil)
	XWhitenT.Product(E, DInv, E.T(), X.T())
	var XWhiten = mat.NewDense(r, c, nil)
	XWhiten = mat.DenseCopyOf(XWhitenT.T())
	return XWhiten, err
}

//CalcNewW calculate new w.
////w is *mat.Dense [1, r]
////X is *mat.Dense [r, c]: [signul num, sample num]
//w is *mat.Dense [c, 1]
//X is *mat.Dense [r, c]: [signul num, sample num]
func CalcNewW(w, X *mat.Dense) *mat.Dense {
	r, c := X.Dims()
	//var auxSl1 = make([]float64, 1*c)
	//var auxSl2 = make([]float64, 1*c)
	//var auxSl3 = make([]float64, 1*c)
	//aux1 := mat.NewDense(1, c, nil)
	var auxSl1 = make([]float64, r)
	var auxSl2 = make([]float64, r)
	var auxSl3 = make([]float64, r)
	//aux1 := mat.NewDense(r, 1, nil)
	aux1 := mat.NewDense(1, r, nil)
	////aux1.Mul(w, X)
	//aux1.Mul(X, w)
	aux1.Mul(w.T(), X.T())
	//row1 := aux1.RawRowView(0)
	//for i := 0; i < c; i++ {
	//	auxSl1[i] = g(row1[i])
	//	auxSl2[i] = gDer(row1[i])
	//}
	col1 := make([]float64, r)
	mat.Col(col1, 0, aux1.T())
	for i := 0; i < c; i++ {
		auxSl1[i] = g(col1[i])
		auxSl2[i] = gDer(col1[i])
	}

	aux3 := mat.NewDense(r, c, nil)
	// diagonal matrix
	//for i := 0; i < r; i++ {
	//	copy(auxSl3, X.RawRowView(i))
	//	floats.Mul(auxSl3, auxSl1)
	//	aux3.SetRow(i, auxSl3)
	//}
	for i := 0; i < c; i++ {
		//copy(auxSl3, X.RawRowView(i))
		mat.Col(auxSl3, i, X)
		floats.Mul(auxSl3, auxSl1)
		aux3.SetCol(i, auxSl3)
	}
	//// [r, 1]
	// [1, c]
	//aux4 := ColMeanVector(aux3)
	aux4 := RowMeanVector(aux3)
	//aux7 := mat.NewDense(1, r, nil)
	//aux7.Scale(SliceMean(auxSl2), w)
	//wNew := mat.NewDense(1, r, nil)
	//wNew.Sub(aux4.T(), aux7)
	aux7 := mat.NewDense(c, 1, nil)
	aux7.Scale(SliceMean(auxSl2), w)
	wNew := mat.NewDense(c, 1, nil)
	wNew.Sub(aux4.T(), aux7)
	aux8 := make([]float64, c)
	//TODO
	//copy(aux8, wNew.RawRowView(0))
	mat.Col(aux8, 0, wNew)
	for i, v := range aux8 {
		aux8[i] = v * v
	}
	wNew.Scale(1/math.Sqrt(floats.Sum(aux8)), wNew)
	return wNew
}

func ICA(X *mat.Dense, iter int, tolerance float64) (*mat.Dense, error) {
	r, c := X.Dims()
	X = center(X)
	X, err := Whiten(X)
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
			// progress
			fmt.Printf("Calculating... %d%%\r", (i*iter+j+1)*100/(componentsNR*iter))

			wNew = CalcNewW(w, X)
			if i >= 1 {
				Wi := W.Slice(0, i, 0, componentsNR)
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
	fmt.Println("\n\ncomplete")
	return S, nil
}
