package adflib

import (
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
			xRow = Unset(xRow, 0)
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
				mu: 1.0,
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

func TestNewFiltFBLMS(t *testing.T) {
	type args struct {
		n  int
		mu float64
		w  interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    FDADFInterface
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFiltFBLMS(tt.args.n, tt.args.mu, tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFiltFBLMS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFiltFBLMS() got = %v, want %v", got, tt.want)
			}
		})
	}
}
