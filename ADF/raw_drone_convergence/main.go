package main

import (
	"bufio"
	"fmt"
	"github.com/tetsuzawa/go-adflib/misc"
	"gonum.org/v1/gonum/floats"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-audio/wav"
	"github.com/tetsuzawa/go-adflib/adf"
)

const (
	eps = 1e-5
)

func main() {
	var (
		err error

		wavName string
		adfName string
		L       int
		mu      float64
		order   int

		dataDir string
	)

	wavName = os.Args[1]
	data := ReadDataFromWav(wavName)

	adfName = os.Args[2]

	L, err = strconv.Atoi(os.Args[3])
	check(err)
	mu, err = strconv.ParseFloat(os.Args[4], 64)
	check(err)
	order, err = strconv.Atoi(os.Args[5])
	check(err)

	w := "zeros"

	var af adf.AdaptiveFilter
	var testName string
	switch adfName {
	case "LMS":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v.csv", adfName, mu, L)
		af, err = adf.NewFiltLMS(L, mu, w)
		check(err)
	case "NLMS":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v.csv", adfName, mu, L)
		af, err = adf.NewFiltNLMS(L, mu, eps, w)
		check(err)
	case "AP":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v_order-%v.csv", adfName, mu, L, order)
		af, err = adf.NewFiltAP(L, mu, order, eps, w)
		check(err)
	case "RLS":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v.csv", adfName, mu, L)
		af, err = adf.NewFiltRLS(L, mu, eps, w)
		check(err)
	}

	dataDir = "data"

	fmt.Println("making d, x ...")
	d, x := makeData(data, L)
	fmt.Println("exploring mu ...")
	mu = ExploreLearning(d, x, af, testName, dataDir)

	fmt.Println("running ...")
	y, e, _, err := af.Run(d, x)
	check(err)

	saveWavAsCSV(d, y, e, dataDir, testName)
}

func ReadDataFromWav(name string) []int {
	f, err := os.Open(name)
	check(err)
	defer f.Close()
	wavFile := wav.NewDecoder(f)
	check(err)

	wavFile.ReadInfo()
	ch := int(wavFile.NumChans)
	//byteRate := int(w.BitDepth/8) * ch
	//bps := byteRate / ch
	fs := int(wavFile.SampleRate)
	fmt.Println("ch", ch, "fs", fs)

	buf, err := wavFile.FullPCMBuffer()
	check(err)

	return buf.Data
}

func makeData(data []int, L int) (d []float64, x [][]float64) {
	n := len(data)
	//input value
	x = make([][]float64, n)
	for i := 0; i < n; i++ {
		x[i] = make([]float64, L)
	}
	//noise
	//desired value
	d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, float64(data[i]))
		copy(x[i], xRow)
		d[i] = float64(data[i])
	}
	return d, x
}

func ExploreLearning(d []float64, x [][]float64, af adf.AdaptiveFilter, dataDir, testName string) float64 {

	es, mus, err := adf.ExploreLearning(af, d, x, 0.00001, 2.0, 100, 0.5, 100, "MSE", nil)
	check(err)

	res := make(map[float64]float64, len(es))
	for i := 0; i < len(es); i++ {
		res[es[i]] = mus[i]
	}
	eMin := floats.Min(es)
	fmt.Printf("the step size mu with the smallest error is %.3f\n", res[eMin])

	fw, err := os.Create(filepath.Join(dataDir, testName+"_opt_mu.log"))
	check(err)
	_, err = fmt.Fprintf(fw, "%g\n", res[eMin])
	check(err)
	err = fw.Close()
	check(err)

	return res[eMin]
}

func saveWavAsCSV(d, y, e []float64, dataDir string, testName string) {
	n := len(d)
	fw, err := os.Create(filepath.Join(dataDir, testName+".csv"))
	check(err)
	writer := bufio.NewWriter(fw)
	for i := 0; i < n; i++ {
		//_, err = fw.Write([]byte(fmt.Sprintf("%f,%f,%f\n", d[i], y[i], e[i])))
		_, err = writer.WriteString(fmt.Sprintf("%g,%g,%g\n", d[i], y[i], e[i]))
		//_, err = fmt.Fprintf(w, "%g,%g,%g\n", d[i], y[i], e[i])
		check(err)
	}
	err = writer.Flush()
	check(err)
	err = fw.Close()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
