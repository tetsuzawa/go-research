package misc

import (
	"errors"
	"github.com/gonum/floats"
	"math"
)

func ElmAbs(fs []float64) []float64 {
	for i, f := range fs {
		fs[i] = math.Abs(f)
	}
	return fs
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
