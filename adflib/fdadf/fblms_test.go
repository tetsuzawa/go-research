package fdadf

import (
	"fmt"
	"github.com/tetsuzawa/go-research/adflib/adf"
	"github.com/tetsuzawa/go-research/adflib/misc"
	"gonum.org/v1/gonum/floats"
	"math/rand"
	"reflect"
	"testing"
)

func TestFiltFBLMS_Run(t *testing.T) {
	rand.Seed(1)
	//creation of data
	//number of samples
	n := 512
	L := 32
	m := n / L
	//input value
	var x = make([][]float64, m)
	//desired value
	var d = make([][]float64, m)
	var xRow = make([]float64, L)
	for i := 0; i < m; i++ {
		for j := 0; j < L; j++ {
			xRow = misc.Unset(xRow, 0)
			xRow = append(xRow, rand.NormFloat64())
		}
		x[i] = append([]float64{}, xRow...)
		copy(d[i], x[i])
	}
	type fields struct {
		n  int
		mu float64
		w  interface{}
	}
	type args struct {
		d [][]float64
		x [][]float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]float64
		want1   [][]float64
		want2   [][]float64
		wantErr bool
	}{
		{
			name: "FBLMS Run",
			fields: fields{
				n:  L,
				mu: 0.0000000001,
				w:  "zeros",
			},
			args: args{
				d: d,
				x: x,
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af, _ := NewFiltFBLMS(tt.fields.n, tt.fields.mu, tt.fields.w)
			got, got1, got2, err := af.Run(tt.args.d, tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Run() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Run() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func ExampleExploreLearning_fblms() {
	rand.Seed(1)
	//creation of data
	//number of samples
	n := 512
	L := 32
	m := n / L
	mu := 0.0000001

	//input value
	var x = make([][]float64, m)
	//desired value
	var d = make([][]float64, m)
	var xRow = make([]float64, L)
	for i := 0; i < m; i++ {
		for j := 0; j < L; j++ {
			xRow = misc.Unset(xRow, 0)
			xRow = append(xRow, rand.NormFloat64())
		}
		x[i] = append([]float64{}, xRow...)
		copy(d[i], x[i])
	}

	af, err := NewFiltFBLMS(L, mu, "zeros")
	check(err)
	es, mus, err := ExploreLearning(af, d, x, 0.00001, 2.0, 100, 0.5, 100, "MSE", nil)
	check(err)

	res := make(map[float64]float64, len(es))
	for i := 0; i < len(es); i++ {
		res[es[i]] = mus[i]
	}
	eMin := floats.Min(es)
	fmt.Printf("the step size mu with the smallest error is %.3f\n", res[eMin])
	//output:
	//the step size mu with the smallest error is 1.313
}
