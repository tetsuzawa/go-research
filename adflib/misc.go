package adflib

import (
	"errors"
	"github.com/gonum/floats"
	"log"
	"math"
	"math/rand"
)

func ElmAbs(fs []float64) []float64 {
	for i, f := range fs {
		fs[i] = math.Abs(f)
	}
	return fs
}

func LogSE(x1, x2 []float64) ([]float64, error) {
	e, err := GetValidError(x1, x2)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(e); i++ {
		e[i] = 10 * math.Log10(math.Pow(e[i], 2))
	}
	return e, nil

}

func MAE(x1, x2 []float64) (float64, error) {
	e, err := GetValidError(x1, x2)
	if err != nil {
		return 0, err
	}
	return floats.Sum(ElmAbs(e)) / float64((len(e))), nil
}

func MSE(x1, x2 []float64) (float64, error) {
	e, err := GetValidError(x1, x2)
	if err != nil {
		return 0, err
	}
	return floats.Dot(e, e) / float64(len(e)), nil
}

func RMSE(x1, x2 []float64) (float64, error) {
	e, err := GetValidError(x1, x2)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(floats.Dot(e, e)) / float64(len(e)), nil
}

func GetValidError(x1, x2 []float64) ([]float64, error) {
	if len(x1) != len(x2) {
		err := errors.New("length of two slices is different")
		return nil, err
	}
	floats.Sub(x1, x2)
	e := x1
	return e, nil
}

func GetMeanError(x1, x2 []float64, fn string) (float64, error) {
	switch fn {
	case "MAE":
		return MAE(x1, x2)
	case "MSE":
		return MSE(x1, x2)
	case "RMSE":
		return RMSE(x1, x2)
	default:
		err := errors.New(`The provided error function (fn) is not known.
								Use "MAE", "MSE" or "RMSE"`)
		return 0, err
	}
}

func Floor(fs [][]float64) []float64 {
	var fs1d = make([]float64, len(fs)*len(fs[0]))
	for i, sl := range fs {
		for j, v := range sl {
			fs1d[len(fs)*i+j] = v
		}
	}
	return fs1d
}


func NewRandSlice(n int) []float64 {
	rs := make([]float64, n)
	for i := 0; i < n; i++ {
		rs[i] = rand.Float64()
	}
	return rs
}

func NewNormRandSlice(n int) []float64 {
	rs := make([]float64, n)
	for i := 0; i < n; i++ {
		rs[i] = rand.NormFloat64()
	}
	return rs
}

// NewRand2dSlice make 2d slice.
// the arg n is sample number and m is number of signals.
func NewRand2dSlice(n, m int) [][]float64 {
	rs2 := make([][]float64, m)
	for j := 0; j < m; j++ {
		rs2[j] = NewRandSlice(n)
	}
	return rs2
}

// NewRandNorm2dSlice make 2d slice.
// the arg n is sample number and m is number of signals.
func NewNormRand2dSlice(n, m int) [][]float64 {
	rs2 := make([][]float64, m)
	for j := 0; j < m; j++ {
		rs2[j] = NewNormRandSlice(n)
	}
	return rs2
}

func Unset(s []float64, i int) []float64 {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
