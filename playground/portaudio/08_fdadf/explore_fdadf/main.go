package main

import (
	"fmt"
	"github.com/gonum/floats"
	"github.com/takuyaohashi/go-wav"
	"github.com/tetsuzawa/converter"
	"github.com/tetsuzawa/go-research/adflib/fdadf"
	"log"
	"os"
	"reflect"
)

const (
	mu              = 0.000000000000000001
	FramesPerBuffer = 1024
	order           = 8
	eps             = 1e-5
)

var af, _ = fdadf.NewFiltFBLMS(FramesPerBuffer, mu, "zeros")

func main() {
	f1, err := os.Open(os.Args[1])
	check(err)
	r1, err := wav.NewReader(f1)
	check(err)
	data1, err := r1.ReadSamples(int(r1.GetSubChunkSize()) / int(r1.GetNumChannels()) / int(r1.GetBitsPerSample()) * 8)
	check(err)

	if reflect.TypeOf(data1) != reflect.TypeOf([]int16{0,}) {
		log.Fatalln(err)
	}

	val1, ok := data1.([]int16)
	if !ok {
		log.Fatalln("Data type is not valid")
	}
	valf1 := converter.Int16sToFloat64s(val1)

	af, err := fdadf.NewFiltFBLMS(FramesPerBuffer, mu, "zeros")
	check(err)
	rvalf1 := reshape(valf1, len(valf1)/FramesPerBuffer, FramesPerBuffer)
	es, mus, err := fdadf.ExploreLearning(af, rvalf1, rvalf1, 0.00001, 1.0, 100, 0.5, 100, "MSE", nil)
	check(err)
	res := make(map[float64]float64, len(es))
	for i := 0; i < len(es); i++ {
		res[es[i]] = mus[i]
	}
	for i := 0; i < len(es); i++ {
		fmt.Println(es[i], mus[i])
	}
	eMin := floats.Min(es)
	fmt.Printf("the step size mu with the smallest error is %.3f\n", res[eMin])

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func reshape(fs []float64, r, c int) [][]float64 {
	if len(fs) != r*c {
		panic(fmt.Sprintf("the length of fs is invalid. got: %d, want: %d", len(fs), r*c))
	}
	var fs2d = make([][]float64, r)
	for i := range fs2d {
		fs2d[i] = make([]float64, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fs2d[i][j] = fs[i*r+j]
		}
	}
	return fs2d
}
