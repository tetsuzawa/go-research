package ica

import (
	"gonum.org/v1/gonum/mat"
	"reflect"
	"testing"
)

func TestColMeanVector(t *testing.T) {
	type args struct {
		X *mat.Dense
	}
	tests := []struct {
		name string
		args args
		want *mat.Dense
	}{
		{
			name: "arange 3x3",
			args: args{X: mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})},
			want: mat.NewDense(3, 1, []float64{2, 5, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColMeanVector(tt.args.X); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColMeanVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRowMeanVector(t *testing.T) {
	type args struct {
		X *mat.Dense
	}
	tests := []struct {
		name string
		args args
		want *mat.Dense
	}{
		{
			name: "arange 3x3",
			args: args{X: mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})},
			want: mat.NewDense(1, 3, []float64{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RowMeanVector(tt.args.X); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RowMeanVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceMean(t *testing.T) {
	type args struct {
		fs []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "123",
			args: args{fs: []float64{1, 2, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceMean(tt.args.fs); got != tt.want {
				t.Errorf("SliceMean() = %v, want %v", got, tt.want)
			}
		})
	}
}
