package adflib

import (
	"math/rand"
	"reflect"
	"testing"
)

func init() {
	rand.Seed(1)
}

func TestAdaptiveFilter_CheckFloatParam(t *testing.T) {
	type fields struct {
		w  []float64
		n  int
		mu float64
	}
	type args struct {
		p    float64
		low  float64
		high float64
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AdaptiveFilter{
				w:  tt.fields.w,
				n:  tt.fields.n,
				mu: tt.fields.mu,
			}
			got, err := af.CheckFloatParam(tt.args.p, tt.args.low, tt.args.high, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckFloatParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckFloatParam() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdaptiveFilter_CheckIntParam(t *testing.T) {
	type fields struct {
		w  []float64
		n  int
		mu float64
	}
	type args struct {
		p    int
		low  int
		high int
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AdaptiveFilter{
				w:  tt.fields.w,
				n:  tt.fields.n,
				mu: tt.fields.mu,
			}
			got, err := af.CheckIntParam(tt.args.p, tt.args.low, tt.args.high, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIntParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckIntParam() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdaptiveFilter_ExploreLearning(t *testing.T) {
	type fields struct {
		w  []float64
		n  int
		mu float64
	}
	type args struct {
		d        []float64
		x        []float64
		muStart  float64
		muEnd    float64
		steps    int
		nTrain   float64
		epochs   int
		criteria string
		targetW  []float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AdaptiveFilter{
				w:  tt.fields.w,
				n:  tt.fields.n,
				mu: tt.fields.mu,
			}
			got, err := af.ExploreLearning(tt.args.d, tt.args.x, tt.args.muStart, tt.args.muEnd, tt.args.steps, tt.args.nTrain, tt.args.epochs, tt.args.criteria, tt.args.targetW)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExploreLearning() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExploreLearning() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdaptiveFilter_InitWeights(t *testing.T) {
	type fields struct {
		w  []float64
		n  int
		mu float64
	}
	type args struct {
		w interface{}
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AdaptiveFilter{
				w:  tt.fields.w,
				n:  tt.fields.n,
				mu: tt.fields.mu,
			}
			if err := af.InitWeights(tt.args.w, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("InitWeights() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdaptiveFilter_PreTrainedRun(t *testing.T) {
	type fields struct {
		w  []float64
		n  int
		mu float64
	}
	type args struct {
		d      []float64
		x      []float64
		nTrain float64
		epochs int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantY  []float64
		wantE  []float64
		wantW  []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AdaptiveFilter{
				w:  tt.fields.w,
				n:  tt.fields.n,
				mu: tt.fields.mu,
			}
			gotY, gotE, gotW := af.PreTrainedRun(tt.args.d, tt.args.x, tt.args.nTrain, tt.args.epochs)
			if !reflect.DeepEqual(gotY, tt.wantY) {
				t.Errorf("PreTrainedRun() gotY = %v, want %v", gotY, tt.wantY)
			}
			if !reflect.DeepEqual(gotE, tt.wantE) {
				t.Errorf("PreTrainedRun() gotE = %v, want %v", gotE, tt.wantE)
			}
			if !reflect.DeepEqual(gotW, tt.wantW) {
				t.Errorf("PreTrainedRun() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestAdaptiveFilter_Predict(t *testing.T) {
	type fields struct {
		w  []float64
		n  int
		mu float64
	}
	type args struct {
		x []float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantY  float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AdaptiveFilter{
				w:  tt.fields.w,
				n:  tt.fields.n,
				mu: tt.fields.mu,
			}
			if gotY := af.Predict(tt.args.x); gotY != tt.wantY {
				t.Errorf("Predict() = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestAdaptiveFilter_Run(t *testing.T) {
	type fields struct {
		w  []float64
		n  int
		mu float64
	}
	type args struct {
		d []float64
		x []float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantY  []float64
		wantE  []float64
		wantW  []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AdaptiveFilter{
				w:  tt.fields.w,
				n:  tt.fields.n,
				mu: tt.fields.mu,
			}
			gotY, gotE, gotW := af.Run(tt.args.d, tt.args.x)
			if !reflect.DeepEqual(gotY, tt.wantY) {
				t.Errorf("Run() gotY = %v, want %v", gotY, tt.wantY)
			}
			if !reflect.DeepEqual(gotE, tt.wantE) {
				t.Errorf("Run() gotE = %v, want %v", gotE, tt.wantE)
			}
			if !reflect.DeepEqual(gotW, tt.wantW) {
				t.Errorf("Run() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_LinSpace(t *testing.T) {
	type args struct {
		start float64
		end   float64
		n     int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			args: args{start: 0, end: 10, n: 21},
			want: []float64{0., 0.5, 1., 1.5, 2., 2.5, 3., 3.5, 4., 4.5, 5.,
				5.5, 6., 6.5, 7., 7.5, 8., 8.5, 9., 9.5, 10.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinSpace(tt.args.start, tt.args.end, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRandn(t *testing.T) {
	type args struct {
		stddev float64
		mean   float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{stddev: 0.5, mean: 0},
			want: -0.6168790887989735,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandn(tt.args.stddev, tt.args.mean); got != tt.want {
				t.Errorf("NewRandn() = %v, want %v", got, tt.want)
			}
		})
	}
}
