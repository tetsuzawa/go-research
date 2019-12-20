package research

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvolve(t *testing.T) {
	type args struct {
		xs []float64
		ys []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "correct",
			args: args{
				xs: []float64{0, 1, 2, 3, 4, 5},
				ys: []float64{0.2, 0.8},
			},
			want: []float64{0, 0.2, 1.2, 2.2, 3.2, 4.2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convolve(tt.args.xs, tt.args.ys); !reflect.DeepEqual(got, tt.want) {
				for i := 0; i < len(got); i++ {
					fmt.Printf("%g, ", got[i])
				}
				fmt.Println("")
				t.Errorf("Convolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
