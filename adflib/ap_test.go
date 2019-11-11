package adflib

import (
	"fmt"
	"github.com/gonum/floats"
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
		x[i] = append([]float64{}, xRow...)
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
			want:    []float64{0, 0, 0, 0.802869025666775, -0.4715562121113624, 0.17713850557648975, 0.025188596209507003, 0.22398088928199358, -1.5036026392534096, 0.6792126657842255, -0.792829023078393, -0.23913973455893855, -0.21133338100272647, 0.6812679878626717, -1.293106506457297, 0.5835957116273562, 0.8327906067969782, -1.2680399401788836, 1.8035073572655014, -1.1942093250292705, 0.4644140508352807, -0.6343970028679353, 1.0418707663353226, -1.2443722002564406, 1.801778777176059, -1.8745456831952088, 0.7435439417743109, -0.5966909911258799, -0.10954253846972284, 0.06401991389533457, 1.189050013577166, -1.8464430409740569, 2.28160173326647, -1.2037527789196263, -0.6122395485979182, 1.3625634611939315, -1.142527457156328, -0.6455597975695448, 1.3154551873271254, -0.555530984785894, 0.42256150201862774, 1.3878968282832123, -2.514359882467028, 2.276377468932429, -2.607252910594373, 1.9912617233245276, -0.362709191811325, -0.6779359647856646, 2.500486168285557, -2.57892788516804, 1.47160711715115, -2.235471567295035, 2.3678429648780877, -1.9269320013894955, 1.27644359679089, -1.70370867268556, 1.3063107540449281, -1.2891134718731783, 0.9480793660710127, 0.0637319361248013, -1.8161477809598745, 1.896482512011969, -2.2521541399671676, 0.8893827150004558,},
			want1:   []float64{0, 0, 0, -2.0366272032647217, -0.049438359041787894, 0.14566674703509014, 0.13361914396692862, -0.9552639054594726, 3.0890066015340327, 0.6196281817332088, 1.5252709488829062, 0.9392606369989234, 1.210959502013989, -0.99780523075676, 2.393835700207318, 0.4061147085811754, -2.2678375289292063, 1.4053903137522343, -2.6496017307210984, 1.3503358203120852, -0.1846703321721092, 1.3371472644301166, -2.111416168688343, 1.5746615858691024, -2.922492869532692, 2.801844153959292, 0.2294904400688743, 1.106807101611896, 0.36108258252887204, 0.102670805022963, -2.8276910914336004, 3.000673854684094, -3.052162681116264, 0.4275395690458592, 2.032519045386798, -1.693020223668901, 0.8954927403260919, 2.3811930542241067, -1.5212891145839853, -0.49613431281805, -1.100085048281767, -3.3495347393681207, 4.500279164448161, -2.3175155770014215, 3.8109845250807903, -2.8298544858438777, -1.0403835365623402, 1.6522191286742862, -4.49344853827704, 2.9601616254679906, -1.0620359225917002, 4.7657100106474415, -3.638085683373581, 1.506304188768172, -1.2742829281882524, 4.086569018117501, -1.1150446362443205, 1.2991196698580172, -1.406392214277252, -0.5401895150909903, 4.547269651984321, -1.8440118733646995, 3.0520305421207334, -0.1640452270567445,},
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
				for i := 0; i < n; i++ {
					fmt.Printf("%g, ", got[i])
				}
				fmt.Println("")
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Run() got1 = %v, want %v", got1, tt.want1)
				for i := 0; i < n; i++ {
					fmt.Printf("%g, ", got1[i])
				}
				fmt.Println("")
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Run() got2 = %v, want %v", got2, tt.want2)
				for i := 0; i < n; i++ {
					fmt.Print("{")
					for k := 0; k < L; k++ {
						fmt.Printf("%g, ", got2[i][k])
					}
					fmt.Print("}, ")
				}
				fmt.Println("")
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

func ExampleExploreLearning_ap() {
	rand.Seed(1)
	//creation of data
	//number of samples
	n := 64
	L := 4
	order := 4
	mu := 1.0
	eps := 0.001
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
		x[i] = append([]float64{}, xRow...)
		v[i] = rand.NormFloat64() * 0.1
		d[i] = x[i][L-1]
	}

	af, err := NewFiltAP(L, mu, order, eps, "random")
	checkError(err)
	es, mus, err := ExploreLearning(af, d, x, 0.0000, 10000.0, 101, 0.5, 100, "MSE", nil)
	checkError(err)

	res := make(map[float64]float64, len(es))
	for i := 0; i < len(es); i++ {
		res[es[i]] = mus[i]
	}
	for i := 0; i < len(es); i++ {
		fmt.Println(es[i], mus[i])
	}
	eMin := floats.Min(es)
	fmt.Printf("the step size mu with the smallest error is %.3f\n", res[eMin])
	//output:
	//the step size mu with the smallest error is 1.313
}
