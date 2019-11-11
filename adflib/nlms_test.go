package adflib

import (
	"fmt"
	"github.com/gonum/floats"
	"math/rand"
	"reflect"
	"testing"
)

func TestFiltNLMS_Run(t *testing.T) {
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
	//	eps            float64
	//	wHistory       [][]float64
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
			name: "Run NLMS Filter",
			args: args{
				d: d,
				x: x,
			},
			want:    []float64{-1.2740625158189787, 1.2740607203730527, -1.274058924929657, 1.2740571294887912, -2.5078117730005216, 1.9868144019668266, -1.6640068043851914, 1.8228119757983932, -2.5540913926709785, 4.139489521454439, -2.840644670812334, 3.5730815613193685, -2.872956610219725, 3.8725772738748514, -4.189108613347012, 5.289830352503853, -4.300113872441337, 2.8650629127734932, -2.7277086952283556, 1.8816116701471697, -1.7254827432604978, 2.00522363610026, -1.3024715390547448, 0.23292580845544283, 0.09736343994963624, -1.2180758157560092, 2.145371263196296, -1.1723352292619051, 1.6824489687885953, -1.430906908251002, 1.5975953757882915, -3.2362318930462592, 4.390456519588648, -5.161010194382629, 4.384790805325514, -2.964507130861138, 2.6340466564058347, -2.881077313132465, 4.616704063784313, -4.822531194980477, 3.7708605833627162, -4.448377860833655, 2.486736445362536, -0.5008164576153015, 0.4596777017542189, 0.7440528641897997, -1.5826433963989353, 0.17955041499725036, 0.7947316289308837, -2.7876900704229626, 3.1689193449833786, -2.759344261870319, 5.289575250989044, -6.559808725198403, 6.139172261064911, -6.137002944007175, 8.51985128300015, -8.32857342831527, 8.338567875331352, -8.796868326717606, 8.32039902238677, -5.589269274790517, 5.641731962933933, -4.84184873749679,},
			want1:   []float64{1.2740625158189787, -1.2740607203730527, 1.274058924929657, -2.507815307086738, 1.9868172018473713, -1.6640091493552467, 1.822814544561627, -2.554094991975872, 4.139495354951602, -2.8406486739370047, 3.573086596616847, -2.872960658879384, 3.8725827312309873, -4.18911451676894, 5.289837807097033, -4.300119932295321, 2.8650669503091084, -2.7277125392001427, 1.8816143217727586, -1.725485174864355, 2.0052264619236695, -1.3024733745380788, 0.23292613670172435, 0.09736357715721891, -1.2180775323062694, 2.1453742865200924, -1.1723368813531105, 1.6824513397479213, -1.430908924729446, 1.5975976271692995, -3.236236453644726, 4.3904627067562965, -5.161017467438442, 4.384796984508862, -2.9645113085366335, 2.6340503683861685, -2.8810813732360705, 4.616710569787027, -4.822537991041173, 3.7708658973765328, -4.448384129625856, 2.4867399497487463, -0.5008171633814036, 0.459678349546309, 0.7440539127321983, -1.5826456267091498, 0.1795506680252701, 0.7947327488913712, -2.787693998922366, 3.168923810722913, -2.759348150423929, 5.289582705222726, -6.559817969484537, 6.139180912577079, -6.137011592462274, 8.519863289439115, -8.328585165199543, 8.33857962630011, -8.79688072353759, 8.320410747751417, -5.589277151362324, 5.641739913437787, -4.841855560780367, 5.567186225440501,},
			want2:   [][]float64{{0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,}, {0.569059219929098, -0.8938905790924524, -0.7964092569819132, -1.6157228347121615,},},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//af := &FiltNLMS{
			//	AdaptiveFilter: tt.fields.AdaptiveFilter,
			//	kind:           tt.fields.kind,
			//	eps:            tt.fields.eps,
			//	wHistory:       tt.fields.wHistory,
			//}
			af := Must(NewFiltNLMS(L, 1.0, 1e-5, "random"))
			got, got1, got2, err := af.Run(tt.args.d, tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v\n, want %v\n", got, tt.want)
				for i := 0; i < n; i++ {
					fmt.Printf("%g, ", got[i])
				}
				fmt.Println("")
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Run() got1 = %v\n, want %v\n", got1, tt.want1)
				for i := 0; i < n; i++ {
					fmt.Printf("%g, ", got1[i])
				}
				fmt.Println("")
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Run() got2 = %v\n, want %v\n", got2, tt.want2)
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
func TestNewFiltNLMS(t *testing.T) {
	type args struct {
		n   int
		mu  float64
		eps float64
		w   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    ADFInterface
		wantErr bool
	}{
		{
			name: "TestNewFiltNLMS",
			args: args{
				n:   4,
				mu:  1.0,
				eps: 1e-5,
				w:   "zeros",
			},
			want: &FiltLMS{
				AdaptiveFilter: AdaptiveFilter{
					w:  mat.NewDense(1, 4, []float64{0, 0, 0, 0}),
					n:  4,
					mu: 1.0,
				},
				kind:     "LMS Filter",
				wHistory: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFiltNLMS(tt.args.n, tt.args.mu, tt.args.eps, tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFiltNLMS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFiltNLMS() got = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

func ExampleExploreLearning_nlms() {
	rand.Seed(1)
	//creation of data
	//number of samples
	//n := 64
	n := 512
	L := 8
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

	af, err := NewFiltNLMS(L, mu, eps, "zeros")
	checkError(err)
	es, mus, err := ExploreLearning(af, d, x, 0.00001, 2.0, 100, 0.5, 100, "MSE", nil)
	checkError(err)

	res := make(map[float64]float64, len(es))
	for i := 0; i < len(es); i++ {
		res[es[i]] = mus[i]
	}
	eMin := floats.Min(es)
	fmt.Printf("the step size mu with the smallest error is %.3f\n", res[eMin])
	//output:
	//the step size mu with the smallest error is 1.313
}
