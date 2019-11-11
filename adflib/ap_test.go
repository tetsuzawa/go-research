package adflib

import (
	"math/rand"
	"reflect"
	"testing"
)

/*
func TestFiltAP_Adapt(t *testing.T) {
	type fields struct {
		AdaptiveFilter AdaptiveFilter
		kind           string
		order          int
		eps            float64
		wHistory       *mat.Dense
		xMem           *mat.Dense
		dMem           *mat.Dense
		yMem           *mat.Dense
		eMem           *mat.Dense
		epsIDE         *mat.Dense
		ide            *mat.Dense
	}
	type args struct {
		d float64
		x []float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &FiltAP{
				AdaptiveFilter: tt.fields.AdaptiveFilter,
				kind:           tt.fields.kind,
				order:          tt.fields.order,
				eps:            tt.fields.eps,
				wHistory:       tt.fields.wHistory,
				xMem:           tt.fields.xMem,
				dMem:           tt.fields.dMem,
				yMem:           tt.fields.yMem,
				eMem:           tt.fields.eMem,
				epsIDE:         tt.fields.epsIDE,
				ide:            tt.fields.ide,
			}
		})
	}
}
*/

func TestFiltAP_Run(t *testing.T) {
	rand.Seed(1)
	//creation of data
	//number of samples
	n := 64
	L := 4
	//input value
	var x = make([][]float64, n)
	//noise
	var v = make([]float64, n)
	//desired value
	var d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = Unset(xRow, 0)
		xRow = append(xRow, rand.NormFloat64())
		x[i] = xRow
		v[i] = rand.NormFloat64() * 0.1
		d[i] = x[i][0]
	}
	//type fields struct {
	//	AdaptiveFilter AdaptiveFilter
	//	kind           string
	//	order          int
	//	eps            float64
	//	wHistory       *mat.Dense
	//	xMem           *mat.Dense
	//	dMem           *mat.Dense
	//	yMem           *mat.Dense
	//	eMem           *mat.Dense
	//	epsIDE         *mat.Dense
	//	ide            *mat.Dense
	//}
	type args struct {
		d []float64
		x [][]float64
	}
	tests := []struct {
		name string
		//fields  fields
		args    args
		want    []float64
		want1   []float64
		want2   [][]float64
		wantErr bool
	}{
		{
			name: "Run AP Filter",
			args: args{
				d: d,
				x: x,
			},
			want:    []float64{-0.47201389445600767, 0.004541467456368742, -0.42434040672862855, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268, -0.19410889600822268,},
			want1:   []float64{0.47201389445600767, -0.004541467456368742, 0.42434040672862855, -1.0396492815897243, -0.3268856751449276, 0.5169141486198026, 0.3529166361846583, -0.5371741201692564, 1.7795128582888458, 1.492949743525657, 0.9265508218127358, 0.8942297984482075, 1.1937350170194851, -0.12242834688586557, 1.2948380897582434, 1.1838193162167543, -1.2409380261240055, 0.33145926958157346, -0.6519854774473745, 0.3502353912910373, 0.47385261467139417, 0.8968591575704039, -0.8754365063447977, 0.5243982816208844, -0.9266051963484105, 1.121407366772306, 1.167143277851408, 0.7042250064942388, 0.4456489400673719, 0.36079961492652024, -1.4445321818482117, 1.34833970971826, -0.5764520518415711, -0.5821043138655444, 1.6143883927971028, -0.1363478664667469, -0.05292582082201336, 1.9297421526627843, -0.011725031248637219, -0.8575564015957213, -0.4834146502549166, -1.7675290150766858, 2.180028177989355, 0.1529707879392302, 1.3978405104946399, -0.6444838665111274, -1.2089838323654425, 1.1683920598968442, -1.7988534739832598, 0.5753426363081732, 0.6036800905676724, 2.7243473393606292, -1.0761338224872705, -0.2265189166131008, 0.19626956461086031, 2.5769692414401635, 0.3853750138088303, 0.2041150939930616, -0.26420395219801657, -0.28234868295796633, 2.925230767032669, 0.24657953465549226, 0.9939852981617883, 0.919446383951934,},
			want2:   [][]float64{{-0.6507507226658572, 0.657011901126719, -0.5912859617953181, 0.31741294852014135,}, {-0.6507507226658572, 0.657011901126719, -0.5912859617953181, 0.31741294852014135,}, {-0.6507507226658572, 0.657011901126719, -0.5912859617953181, 0.31741294852014135,}, {-0.6507507226658572, 0.657011901126719, -0.5912859617953181, 0.31741294852014135,},},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//af := &FiltAP{
			//	AdaptiveFilter: tt.fields.AdaptiveFilter,
			//	kind:           tt.fields.kind,
			//	order:          tt.fields.order,
			//	eps:            tt.fields.eps,
			//	wHistory:       tt.fields.wHistory,
			//	xMem:           tt.fields.xMem,
			//	dMem:           tt.fields.dMem,
			//	yMem:           tt.fields.yMem,
			//	eMem:           tt.fields.eMem,
			//	epsIDE:         tt.fields.epsIDE,
			//	ide:            tt.fields.ide,
			//}
			af := Must(NewFiltAP(L, 1.0, 4, 1e-5, "random"))
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

/*
func TestNewFiltAP(t *testing.T) {
	type args struct {
		n     int
		mu    float64
		order int
		eps   float64
		w     interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    ADFInterface
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFiltAP(tt.args.n, tt.args.mu, tt.args.order, tt.args.eps, tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFiltAP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFiltAP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
