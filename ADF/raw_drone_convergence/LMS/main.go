package main

import (
	"fmt"
	"github.com/tetsuzawa/go-adflib/adf"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/go-audio/wav"
	"github.com/tetsuzawa/go-adflib/misc"
)

const (
	eps = 1e-5
)

var (
	af  adf.ADFInterface
	err error

	wavName string
	adfName string
	L       int
	mu      float64
	order   int
)

func main() {


	wavName = os.Args[1]
	data := openWav(wavName)

	adfName = os.Args[2]

	L, err = strconv.Atoi(os.Args[3])
	check(err)
	mu, err = strconv.ParseFloat(os.Args[4], 64)
	check(err)
	order, err = strconv.Atoi(os.Args[5])
	check(err)

	w := "zeros"

	switch adfName {
	case "LMS":
		af, err = adf.NewFiltLMS(L, mu, w)
		check(err)
	case "NLMS":
		af, err = adf.NewFiltNLMS(L, mu, eps, w)
		check(err)
	case "AP":
		af, err = adf.NewFiltAP(L, mu, order, eps, w)
		check(err)
	case "RLS":
		af, err = adf.NewFiltRLS(L, mu, eps, w)
		check(err)
	}
	run(data, af)

}

func openWav(name string) []int16 {
	f, err := os.Open(name + ".wav")
	check(err)
	defer f.Close()
	w := wav.NewDecoder(f)
	check(err)

	w.ReadInfo()
	ch := int(w.NumChans)
	//byteRate := int(w.BitDepth/8) * ch
	//bps := byteRate / ch
	fs := int(w.SampleRate)
	fmt.Println("ch", ch, "fs", fs)

	aBuf, err := w.FullPCMBuffer()
	check(err)

	return aBuf.Data
}

func run(data []int16,  af adf.ADFInterface) {
	//creation of data

	n := len(data)
	L := 64
	//input value
	var x = make([][]float64, n)
	for i := 0; i < n; i++ {
		x[i] = make([]float64, i)
	}
	//noise
	//desired value
	var d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, float64(data[i]))
		copy(x[i], xRow)
		d[i] = float64(data[i])
	}

	y, e, w, err := af.Run(d, x)
	check(err)

	name := fmt.Sprintf("lms_ex_mu-%v_L-%v.png", mu, L)
	fw, err := os.Create(name)
	if err != nil{
		log.Fatalln(err)
	}
	defer fw.Close()
	for i:=0; i<n; i++ {
		_, err = fw.Write([]byte(fmt.Sprintf("%f,%f,%f\n", d[i], y[i], e[i])))
		if err != nil{
			log.Fatalln(err)
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}

}

func ExampleExploreLearning_lms() {
	rand.Seed(1)
	//creation of data
	//number of samples
	n := 64
	L := 4
	mu := 0.1
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

	af, err := NewFiltLMS(L, mu, "zeros")
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
	//the step size mu with the smallest error is 0.182
}
