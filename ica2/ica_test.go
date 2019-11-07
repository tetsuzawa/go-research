package ica

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"reflect"
	"testing"
)

func init() {
	rand.Seed(1)
}

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
		X *mat.Dense
	}
	tests := []struct {
		name string
		args args
		want *mat.Dense
	}{
		{
			name: "same as python ica2",
			args: args{
				X: mat.NewDense(6, 3, []float64{
					-1.0, -1.0, -2.0,
					-0.8583741434275802, -1.82918707171379, -0.6875612151413701,
					-1.4834507951495062, -2.541725397574753, -2.0251761927242597,
					1.4256732187770198, 2.5128366093885104, 1.9385098281655302,
					1.031509825101539, 1.9157549125507696, 0.9472647376523085,
					-2.2879033166650653, -3.143951658332533, -3.4318549749975977,
				})},
			want: mat.NewDense(6, 3, []float64{
				-0.47124246477273457, -0.3189545657197007, -1.1235303638257683,
				-0.3296166082003148, -1.1481416374334907, 0.18890842103286143,
				-0.9546932599222407, -1.8606799632944537, -1.148706556550028,
				1.9544307540042851, 3.19388204366881, 2.814979464339762,
				1.5602673603288046, 2.5968003468310688, 1.82373437382654,
				-1.7591457814378, -2.4629062240522335, -2.555385338823366,
			}),
		},

		//{name: "arange 1to9", args: args{X: mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})}, want: mat.NewDense(3, 1, []float64{2, 5, 8})},
		//{name: "arange 1to9 with minus", args: args{X: mat.NewDense(3, 3, []float64{1, -2, -3, 4, 5, -6, 7, -8, 9})}, want: mat.NewDense(3, 1, []float64{-4. / 3., 1., 8. / 3.})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := center(tt.args.X); !reflect.DeepEqual(got, tt.want) {
				//fgot := mat.Formatted(got, mat.Prefix(""), mat.Squeeze())
				//fwant := mat.Formatted(tt.want, mat.Prefix(""), mat.Squeeze())
				//t.Errorf("center() = %v, want %v", fgot, fwant)
				t.Errorf("CalcNewW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whiten(t *testing.T) {
	type args struct {
		X *mat.Dense
	}
	tests := []struct {
		name    string
		args    args
		want    *mat.Dense
		wantErr bool
	}{
		{
			name:    "same as python ica2",
			args:    args{X: mat.NewDense(6, 3, []float64{-0.47124246477273457, -0.3189545657197007, -1.1235303638257683, -0.3296166082003148, -1.1481416374334907, 0.18890842103286143, -0.9546932599222407, -1.8606799632944537, -1.148706556550028, 1.9544307540042851, 3.19388204366881, 2.814979464339762, 1.5602673603288046, 2.5968003468310688, 1.82373437382654, -1.7591457814378, -2.4629062240522335, -2.555385338823366,})},
			want:    mat.NewDense(6, 3, []float64{0.07346330684488667, 0.4698717643163002, -1.154327000226922, -0.16266386346927497, -1.1824477251367067, 1.0997098202708528, 0.6183995645352338, -1.17260030854777, -0.2661049624845879, -0.49693713464409583, 1.2247825958531244, 1.2893650136690384, 1.4566654950959261, 0.7096003882016977, -0.1267190081732057, -1.488927368362683, -0.049206714686642616, -0.8419238630551718,}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Whiten(tt.args.X)
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

func Test_CalcNewW(t *testing.T) {
	type args struct {
		w *mat.Dense
		X *mat.Dense
	}
	tests := []struct {
		name string
		args args
		want *mat.Dense
	}{
		{
			name: "same as python ica2",
			args: args{
				w: mat.NewDense(1, 3, []float64{0.5488135039273248, 0.7151893663724195, 0.6027633760716439}),
				X: mat.NewDense(3, 6, []float64{0.07346330684499848, -0.16266386346933273, 0.6183995645355771, -0.4969371346445325, 1.456665495096312, -1.4889273683630337, 0.4698717643162649, -1.1824477251366896, -1.1726003085478844, 1.224782595853268, 0.7096003882015676, -0.049206714686523594, -1.1543270002269619, 1.0997098202708755, -0.266104962484705, 1.2893650136691885, -0.12671900817333798, -0.8419238630550545}),
			},
			want: mat.NewDense(1, 3, []float64{-0.09450675071422016, 0.9885266335381587, 0.11782855704435635}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcNewW(tt.args.w, tt.args.X); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcNewW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestICA(t *testing.T) {
	type args struct {
		X         *mat.Dense
		iter      int
		tolerance float64
	}
	tests := []struct {
		name    string
		args    args
		want    *mat.Dense
		wantErr bool
	}{
		{
			name: "same as python ica2",
			args: args{
				X: mat.NewDense(3, 6, []float64{
					-1.0, -0.8583741434275802, -1.4834507951495062, 1.4256732187770198, 1.031509825101539, -2.2879033166650653,
					-1.0, -1.82918707171379, -2.541725397574753, 2.5128366093885104, 1.9157549125507696, -3.143951658332533,
					-2.0, -0.6875612151413701, -2.0251761927242597, 1.9385098281655302, 0.9472647376523085, -3.4318549749975977,
				}),
				iter:      1000,
				tolerance: 1e-5},
			want: mat.NewDense(3, 6, []float64{
				0.37859833427285494, 0.1446837499030288, 1.2148783695048704, -1.8316384530497773, 0.030429949511564677, 0.06304804985745188,
				1.1366705660266465, -1.605030920865302, -0.46541169689423, -0.17861359965715773, 0.9194395186602825, 0.19294613272975847,
				-0.3403533516886723, 0.18545183406168236, 0.35731399511731016, 0.1461487307780256, 1.2998905058924537, -1.6484517141608042,
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ICA(tt.args.X, tt.args.iter, tt.args.tolerance)
			if (err != nil) != tt.wantErr {
				t.Errorf("ICA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ICA() got = %v, want %v", got, tt.want)
			}
		})
	}
}
