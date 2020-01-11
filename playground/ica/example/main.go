package main

import (
	"errors"
	"github.com/takuyaohashi/go-wav"
	"github.com/tetsuzawa/converter"
	"github.com/tetsuzawa/go-research/ica"
	"log"
	"os"
	"reflect"
)

func main() {
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
