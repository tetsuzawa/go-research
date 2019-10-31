package main

import (
	"fmt"
	"github.com/tetsuzawa/go-research/adflib"
	"log"
	"math/rand"
	"os"
)

func init() {
	rand.Seed(1)
}

func unset(s []float64, i int) []float64 {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

const (
	//step size of filter
	mu = 0.1
	//length of filter
	L = 64
)

func main() {
	//creation of data
	//number of samples
	n := 512
	//input value
	var x = make([][]float64, n)
	//noise
	var v = make([]float64, n)
	//desired value
	var d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = unset(xRow, 0)
		xRow = append(xRow, rand.NormFloat64())
		x[i] = xRow
		v[i] = rand.NormFloat64() * 0.1
		d[i] = x[i][0]
	}

	//identification
	f, err := adflib.NewFiltLMS(L, mu, "zeros")
	if err != nil {
		log.Fatalln(err)
	}

	f.

	y, e, _, err := f.Run(d, x)
	if err != nil {
		log.Fatalln(err)
	}


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
