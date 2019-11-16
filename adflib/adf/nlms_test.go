package adf

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/gonum/floats"
	"github.com/tetsuzawa/go-research/adflib/misc"
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
		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, rand.NormFloat64())
		x[i] = append([]float64{}, xRow...)
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
			want:    []float64{-0.3916108209122005, 0.16536933657985348, -0.0437579947155992, 0.012124947341295034, -0.2654555121700104, 0.3802942691699313, -0.015105143735420865, 0.06338390297796329, -0.35634062205002437, 1.5687950887982662, -0.13433136700109533, 0.2233351641247334, 0.19138298538219461, -0.24457092974407102, 0.014496883486147707, -0.16389871128240147, -0.5703468062521424, 0.12325158184348353, -0.0001354304685272777, 0.41477231714128937, 0.1284912512030528, -0.10594739642904036, -0.6771124260933329, 0.10022071550746274, -0.0007792286060597231, -0.16438259032950334, 0.804033784658881, 0.05122531204556498, -0.3013910495723021, -0.405364406872704, -0.3028385843710033, 0.8416786588506451, -0.1280995968578023, 0.1724792387475776, 0.6525066762526404, -0.17330306888775004, 0.023775619911505473, -0.009171657096201392, 0.7178877185849712, 0.2772808354985652, 0.27493273050125877, 0.31136108806500473, 1.238565646136781, 0.004408117980679305, 0.013802766783230742, -0.7373018741864878, 0.03857756458510833, 0.801736882960992, -0.049745493036917665, 0.5057272250487606, 0.022230842171991626, -0.07618246803444377, -0.8511889527808415, -0.05578977998644839, -0.028604330951522903, 0.0024669022429527057, 0.2697704378237238, 0.17988582851926682, 0.02031316996455955, 0.05439092562840408, 0.034151431454050384, -0.01906135272886425, 0.027614795022829132, 0.6631737158019635,},
			want1:   []float64{0.3916108209122005, -0.16536933657985348, 0.0437579947155992, -1.245883124939242, -0.2555390589831399, -0.057489016558351425, 0.1739128839118565, -0.7946669191554423, 1.9417445843306473, -0.269954241280832, 0.8667732928056084, 0.4767857383152514, 0.8082431356290678, -0.07196631315001722, 1.0862323102638731, 1.1536091314909331, -0.8647001158800858, 0.01409879172986725, -0.8459589429870699, -0.25864582185847473, 0.1512524674601187, 0.8086976579912216, -0.39243297625968754, 0.230068670105199, -1.1199348637505735, 1.0916810610935865, 0.16900059718430427, 0.45889079844045116, 0.5529310936314513, 0.5720551257910016, -1.3358024934854311, 0.31255215485939214, -0.6424613509919914, -0.9486924486213446, 0.7677728205362397, -0.15715369358721953, -0.2708103367417415, 1.744804913750763, -0.9237216458418311, -1.3289461331025092, -0.952456276764398, -2.2729989991499133, 0.7473536358443511, -0.04554622604967179, 1.1899288477031864, -0.10129088833286226, -1.4416702929587735, 0.17254628092762958, -1.9432168769545648, -0.12449348474881006, 0.38734035238745806, 2.6064209113868504, -0.4190537657146517, -0.3648380326348751, 0.03076499955416054, 2.380393443188988, -0.07850432002311614, -0.1698796305344279, -0.4786260181707988, -0.5308485045945931, 2.696970439570396, 0.07153199137613384, 0.7722616071307364, 0.06216377214174784,},
			want2:   [][]float64{{0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,}, {0.006354173589735487, -0.00998127384789638, -0.008892787411415373, -0.018041326816459872,},},
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
		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, rand.NormFloat64())
		x[i] = append([]float64{}, xRow...)
		v[i] = rand.NormFloat64() * 0.1
		d[i] = x[i][L-1]
	}

	af, err := NewFiltNLMS(L, mu, eps, "zeros")
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