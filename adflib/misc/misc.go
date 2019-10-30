package misc

import (
	"errors"
	"github.com/gonum/floats"
	"math"
)

func ElmAbs(fs []float64) []float64 {
	for i, f := range fs{
		fs[i] = math.Abs(f)
	}
	return fs
}

func MAE(x1, x2 []float64) float64 {
	e := GetValidError()
	return floats.Sum(ElmAbs(e)) / float64((len(e)))
}

func MSE(x1, x2 []float64) float64 {
	e := GetValidError()
	return floats.Dot(e, e) / float64(len(e))
}

func RMSE(x1, x2 []float64) float64 {
	e := GetValidError()
	return math.Sqrt(floats.Dot(e, e)) / float64(len(e))
}

func GetValidError(x1, x2 []float64) ([]float64, error) {
	if len(x1) != len(x2){
		err := errors.New("length of slices is different from n")
		return nil,
	}
}

func GetMeanError(x1, x2 float64, fn string) {

}
