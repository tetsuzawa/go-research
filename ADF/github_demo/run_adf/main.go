package main

import (
	"fmt"
	"github.com/tetsuzawa/go-adflib/misc"
	"math"
	"math/rand"
	"os"

	"github.com/tetsuzawa/go-adflib/adf"
	research "github.com/tetsuzawa/go-research/ADF/raw_drone_convergence"
)

func main() {
	dataDir := os.Args[1]

	rand.Seed(2)
	//creation of data
	//number of samples
	n := 256
	L := 8
	mu := 1.0
	//input value
	var x = make([][]float64, n)
	for i := 0; i < n; i++ {
		x[i] = make([]float64, L)
	}
	//noise
	var v = make([]float64, n)
	//desired value
	var d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, 0.2*rand.NormFloat64()+math.Sin(2*math.Pi*1200*float64(i)/48000))
		copy(x[i], xRow)
		v[i] = rand.NormFloat64() * 0.1
		d[i] = x[i][0]
	}
	af := adf.Must(adf.NewFiltNLMS(L, mu, 1e-5, nil))
	testName := fmt.Sprintf("rand_L-%v_mu-%v", L, mu)

	fmt.Println("running ...")
	y, e, _, err := af.Run(d, x)
	check(err)
	research.SaveFilterdDataAsCSV(d, y, e, dataDir, testName)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
