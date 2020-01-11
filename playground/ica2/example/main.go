package main

import (
	"errors"
	"github.com/takuyaohashi/go-wav"
	"github.com/tetsuzawa/converter"
	"github.com/tetsuzawa/go-research/ica2"
	"gonum.org/v1/gonum/mat"
	"log"
	"os"
	"reflect"
)

func main() {
	//calcNewWEx()
	//ExICA()
	m2()
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

func whitenEx() {
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
		_, err := ica.Whiten(tt.args.x)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
func ExICA() {
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
				-0.1195232895154406, -0.13695939213469, 0.25336466042164585, 0.13255584827759057, 1.4968211161532372, -1.6262589432023518,
				0.37608121950301976, 0.1446355752912923, 1.218168432335386, -1.829777924611693, 0.046902947757951625, 0.04398974972403464,
				-1.1448585547570456, 1.5568784779575202, 0.5111648717750966, 0.20006227927267833, -0.610400384720045, -0.5128466895282026,
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := ica.ICA(tt.args.X, tt.args.iter, tt.args.tolerance)
		if (err != nil) != tt.wantErr {
			log.Fatalf("ICA() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			log.Fatalf("ICA() got = %v, want %v", got, tt.want)
		}
	}
}
func m2() {
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

	//data := make([][]float64, 3)
	//data[0] = valf1
	//data[1] = valf2
	//data[2] = valf3

	//ICA := ica.NewICA(data)
	//y, err := ICA.CalcICA()
	var allData []float64
	allData = append(allData, valf1...)
	allData = append(allData, valf2...)
	allData = append(allData, valf3...)

	X := mat.NewDense(3, len(valf1), allData)
	//yMat, err := ica.ICA(X, 1000, 1e-5)
	yMat, err := ica.ICA(X, 1000, 1.7)
	var y =make([][]float64, 3)
	y[0] = yMat.RawRowView(0)
	y[1] = yMat.RawRowView(1)
	y[2] = yMat.RawRowView(2)


	p := wav.WriterParam{
		SampleRate:    44100,
		BitsPerSample: 16,
		NumChannels:   1,
		AudioFormat:   1,
	}

	fw1, err := os.Create("out_1.wav")
	if err != nil {
		log.Fatalln(err)
	}
	fw2, err := os.Create("out_2.wav")
	if err != nil {
		log.Fatalln(err)
	}
	fw3, err := os.Create("out_3.wav")
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

/*
func m1() {
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
