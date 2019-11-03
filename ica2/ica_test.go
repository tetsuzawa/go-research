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
		{name: "plus 3", args: args{x: 3.}, want: 0.9950547536867305,},
		{name: "plus -2", args: args{x: -2.}, want: -0.9640275800758169,},
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
		{name: "plus 3", args: args{x: 3.}, want: 0.009866037165440211,},
		{name: "plus -2", args: args{x: -2.}, want: 0.07065082485316443,},
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
			args:    args{x: mat.NewDense(3, 6, []float64{-0.47124246477273457, -0.3296166082003148, -0.9546932599222407, 1.9544307540042851, 1.5602673603288046, -1.7591457814378, -0.3189545657197007, -1.1481416374334907, -1.8606799632944537, 3.19388204366881, 2.5968003468310688, -2.4629062240522335, -1.1235303638257683, 0.18890842103286143, -1.148706556550028, 2.814979464339762, 1.82373437382654, -2.555385338823366,})},
			want:    mat.NewDense(3, 6, []float64{0.07346330684488667, - 0.16266386346927497, 0.6183995645352338, -0.49693713464409583, 1.4566654950959261, -1.488927368362683, 0.4698717643163002, -1.1824477251367067, -1.17260030854777, 1.2247825958531244, 0.7096003882016977, -0.049206714686642616, -1.154327000226922, 1.0997098202708528, -0.2661049624845879, 1.2893650136690384, -0.1267190081732057, -0.8419238630551718,}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Whitening(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Whitening() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				//fgot := mat.Formatted(got, mat.Prefix(""), mat.Squeeze())
				//fwant := mat.Formatted(tt.want, mat.Prefix(""), mat.Squeeze())
				t.Errorf("Whitening() got = %v, want %v", got, tt.want)
				//t.Errorf("Whitening() got = \n%v\n, want \n%v\n", fgot, fwant)
			}
		})
	}
}
