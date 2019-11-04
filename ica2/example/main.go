package main

import (
	"github.com/tetsuzawa/go-research/ica2"
	"gonum.org/v1/gonum/mat"
	"log"
	"reflect"
)

func main() {
	calcNewWEx()
}

func calcNewWEx() {
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
			want: mat.NewDense(1, 3, []float64{0.0003703279502838618, 0.040517027640052605, 0.0005756545563130041}),
		},
	}
	for _, tt := range tests {
		if got := ica.CalcNewW(tt.args.w, tt.args.X); !reflect.DeepEqual(got, tt.want) {
			log.Fatalf("calcNewW() = %v, want %v", got, tt.want)
		}
	}
}

func whiteningEx() {
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
			want:    mat.NewDense(3, 6, []float64{0.07346330684499848, -0.16266386346933273, 0.6183995645355771, -0.4969371346445325, 1.456665495096312, -1.4889273683630337, 0.4698717643162649, -1.1824477251366896, -1.1726003085478844, 1.224782595853268, 0.7096003882015676, -0.049206714686523594, -1.1543270002269619, 1.0997098202708755, -0.266104962484705, 1.2893650136691885, -0.12671900817333798, -0.8419238630550545,}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		_, err := ica.Whitening(tt.args.x)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

/*
func m() {
	f1, err := os.Open("mix_1.wav")
	if err != nil {
		log.Fatalln(err)
	}
	f2, err := os.Open("mix_2.wav")
	if err != nil {
		log.Fatalln(err)
	}
	f3, err := os.Open("mix_3.wav")
	if err != nil {
		log.Fatalln(err)
	}

	r1, err := wav.NewReader(f1)
	r2, err := wav.NewReader(f2)
	r3, err := wav.NewReader(f3)
	data1, err := r1.ReadSamples(int(r1.GetSubChunkSize()) / int(r1.GetNumChannels()) / int(r1.GetBitsPerSample()) * 8)
	if err != nil {
		log.Fatalln(err)
	}

	if reflect.TypeOf(data1) != reflect.TypeOf([]int16{0,}) {
		log.Fatalln(err)
	}

	val1, ok := data1.([]int16)
	if !ok {
		log.Fatalln("Data type is not valid")
	}
	data2, err := r2.ReadSamples(int(r1.GetSubChunkSize()) / int(r1.GetNumChannels()) / int(r1.GetBitsPerSample()) * 8)
	if err != nil {
		log.Fatalln(err)
	}

	if reflect.TypeOf(data2) != reflect.TypeOf([]int16{0,}) {
		log.Fatalln(err)
	}

	val2, ok := data2.([]int16)
	if !ok {
		log.Fatalln("Data type is not valid")
	}
	data3, err := r3.ReadSamples(int(r1.GetSubChunkSize()) / int(r1.GetNumChannels()) / int(r1.GetBitsPerSample()) * 8)
	if err != nil {
		log.Fatalln(err)
	}

	if reflect.TypeOf(data3) != reflect.TypeOf([]int16{0,}) {
		log.Fatalln(err)
	}

	val3, ok := data3.([]int16)
	if !ok {
		log.Fatalln("Data type is not valid")
	}
	if len(val1) != len(val2) || len(val1) != len(val3) {
		log.Fatalln(errors.New("length is not agree"))
	}

	valf1 := converter.Int16sToFloat64s(val1)
	valf2 := converter.Int16sToFloat64s(val2)
	valf3 := converter.Int16sToFloat64s(val3)

	data := make([][]float64, 3)
	data[0] = valf1
	data[1] = valf2
	data[2] = valf3

	ICA := ica.NewICA(data)
	y, err := ICA.CalcICA()

	p := wav.WriterParam{
		SampleRate:    44100,
		BitsPerSample: 16,
		NumChannels:   1,
		AudioFormat:   1,
	}

	fw1, err := os.Open("out_1.wav")
	if err != nil {
		log.Fatalln(err)
	}
	fw2, err := os.Open("out_2.wav")
	if err != nil {
		log.Fatalln(err)
	}
	fw3, err := os.Open("out_3.wav")
	if err != nil {
		log.Fatalln(err)
	}
	w1, err := wav.NewWriter(fw1, p)
	if err != nil {
		log.Fatalln(err)
	}
	defer w1.Close()
	w2, err := wav.NewWriter(fw2, p)
	if err != nil {
		log.Fatalln(err)
	}
	defer w2.Close()
	w3, err := wav.NewWriter(fw3, p)
	if err != nil {
		log.Fatalln(err)
	}
	defer w3.Close()

	w1.WriteSamples(converter.Float64sToInt16s(y[0]))
	w2.WriteSamples(converter.Float64sToInt16s(y[1]))
	w3.WriteSamples(converter.Float64sToInt16s(y[2]))

}

*/
