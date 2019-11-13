package adflib

import (
	"github.com/gonum/floats"
	"github.com/mjibson/go-dsp/fft"
	"github.com/pkg/errors"
	"github.com/tetsuzawa/converter"
	"math/cmplx"
)

type FiltFBLMS struct {
	AdaptiveFilter
	w        []float64
	kind     string
	wHistory [][]float64
	xMem     []float64
}

func NewFiltFBLMS(n int, mu float64, w interface{}) (FDADFInterface, error) {
	var err error
	p := new(FiltFBLMS)
	p.kind = "FBLMS filter"
	p.n = n
	p.mu, err = p.CheckFloatParam(mu, 0, 1000, "mu")
	if err != nil {
		return nil, errors.Wrap(err, "Parameter error at CheckFloatParam()")
	}
	err = p.InitWeights(w, 2*n)
	if err != nil {
		return nil, err
	}
	p.xMem = make([]float64, 2*n)
	return p, nil
}

func (af *FiltFBLMS) Adapt(d float64, x []float64) {
	zeros := make([]float64, af.n)
	Y := make([]complex128, 2*af.n)
	y := make([]float64, af.n)
	e := make([]float64, af.n)
	EU := make([]complex128, 2*af.n)

	W := fft.FFT(converter.Float64sToComplex128s(append(af.w[:af.n], zeros...)))
	U := fft.FFT(converter.Float64sToComplex128s(append(af.xMem, x...)))
	for i := 0; i < 2*af.n; i++ {
		Y[i] = W[i] * U[i]
	}
	yc := fft.IFFT(Y)[af.n:]
	for i := 0; i < 2; i++ {
		y[i] = real(yc[i])
		e[i] = x[i] - y[i]
	}

	// 2 compute the correlation vector
	aux1 := fft.FFT(converter.Float64sToComplex128s(append(zeros, e...)))
	aux2 := fft.FFT(converter.Float64sToComplex128s(x))
	for i := 0; i < 2*af.n; i++ {
		EU[i] = aux1[i] * cmplx.Conj(aux2[i])
	}
	phi := fft.IFFT(EU)[:af.n]

	// 3 update the parameters of the filter
	aux1 = fft.FFT(converter.Float64sToComplex128s(append(af.w[:af.n], zeros...)))
	aux2 = fft.FFT(append(phi, converter.Float64sToComplex128s(zeros)...))
	for i := 0; i < 2*af.n; i++ {
		W[i] = aux1[i] + complex(af.mu, 0)*aux2[i]
	}
	aux3 := fft.IFFT(W)
	for i := 0; i < 2*af.n; i++ {
		af.w[i] = real(aux3[i])
	}

}

func (af *FiltFBLMS) Predict(x []float64) (y []float64) {
	zeros := make([]float64, af.n)
	y = make([]float64, af.n)
	Y := make([]complex128, 2*af.n)
	W := fft.FFT(converter.Float64sToComplex128s(append(af.w[:af.n], zeros...)))
	U := fft.FFT(converter.Float64sToComplex128s(append(af.xMem, x...)))
	for i := 0; i < 2*af.n; i++ {
		Y[i] = W[i] * U[i]
	}
	yc := fft.IFFT(Y)[af.n:]
	for i := 0; i < 2; i++ {
		y[i] = real(yc[i])
	}
	return
}

func (af *FiltFBLMS) Run(d []float64, x [][]float64) ([]float64, []float64, [][]float64, error) {
	//measure the data and check if the dimension agree
	N := len(x)
	if len(d) != N {
		return nil, nil, nil, errors.New("the length of slice d and x must agree")
	}
	af.n = len(x[0])
	af.wHistory = make([][]float64, N)

	y := make([]float64, N)
	e := make([]float64, N)
	//adaptation loop
	for i := 0; i < N; i++ {
		w := af.w.RawRowView(0)
		af.wHistory[i] = w
		y[i] = floats.Dot(w, x[i])
		e[i] = d[i] - y[i]
		for j := 0; j < af.n; j++ {
			w[j] = af.mu * e[i] * x[i][j]
		}
	}
	return y, e, af.wHistory, nil
}
