package ica

import (
	"gonum.org/v1/gonum/mat"
	"reflect"
	"testing"
)

func Test_g(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "plus 3", args: args{x: 3.}, want: 0.995055,},
		{name: "plus -2", args: args{x: -2.}, want: -0.964028,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g(tt.args.x); got != tt.want {
				t.Errorf("g() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gDer(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "plus 3", args: args{x: 3.}, want: 0.009866,},
		{name: "plus -2", args: args{x: -2.}, want: 0.070651,},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gDer(tt.args.x); got != tt.want {
				t.Errorf("gDer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_center(t *testing.T) {
	type args struct {
		x *mat.Dense
	}
	tests := []struct {
		name string
		args args
		want *mat.Dense
	}{
		{name: "arange 1to9", args: args{x: mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})}, want: mat.NewDense(3, 1, []float64{2, 5, 8})},
		{name: "arange 1to9 with minus", args: args{x: mat.NewDense(3, 3, []float64{1, -2, -3, 4, 5, -6, 7, -8, 9})}, want: mat.NewDense(3, 1, []float64{-4. / 3., 1., 8. / 3.})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := center(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				fgot := mat.Formatted(got, mat.Prefix(""), mat.Squeeze())
				fwant := mat.Formatted(tt.want, mat.Prefix(""), mat.Squeeze())
				t.Errorf("center() = %v, want %v", fgot, fwant)
			}
		})
	}
}

func Test_whitening(t *testing.T) {
	type args struct {
		x *mat.Dense
	}
	tests := []struct {
		name    string
		args    args
		want    *mat.Dense
		wantErr bool
	}{
		{
			name:    "same as python ica2",
			args:    args{},
			want:    nil,
			wantErr: false,
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := whitening(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("whitening() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whitening() got = %v, want %v", got, tt.want)
			}
		})
	}
}