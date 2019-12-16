package main

import (
	"encoding/json"
	"fmt"
	"github.com/tetsuzawa/go-adflib/misc"
	"io/ioutil"
	"math"
	"math/rand"
	"os"

	"github.com/tetsuzawa/go-adflib/adf"
	research "github.com/tetsuzawa/go-research/ADF/raw_drone_convergence"
)

const (
	eps             = 1e-5
	applicationName = "auto_on"
)

func main() {

	jsonName := os.Args[1]
	dataDir := os.Args[2]

	rawJSON, err := ioutil.ReadFile(jsonName)
	check(err)

	var optStepADF = new(research.OptStepADF)
	err = json.Unmarshal(rawJSON, optStepADF)
	check(err)

	wavName := optStepADF.WavName
	adfName := optStepADF.AdfName

	data := research.ReadDataFromWav(wavName)
	L := optStepADF.L
	mu := optStepADF.Mu
	order := optStepADF.Order
	fmt.Println("wav name:", wavName)
	fmt.Println("adf name:", adfName)
	fmt.Println("L:", L)
	fmt.Println("mu:", mu)
	if adfName == "AP" {
		fmt.Println("order:", order)
	}

	var af adf.AdaptiveFilter
	var testName string
	switch adfName {
	case "LMS":
		testName = fmt.Sprintf("%v_%v_L-%v", adfName, applicationName, L)
		af, err = adf.NewFiltLMS(L, mu, nil)
		check(err)
	case "NLMS":
		testName = fmt.Sprintf("%v_%v_L-%v", adfName, applicationName, L)
		af, err = adf.NewFiltNLMS(L, mu, eps, nil)
		check(err)
	case "AP":
		testName = fmt.Sprintf("%v_%v_L-%v_order-%v", adfName, applicationName, L, order)
		af, err = adf.NewFiltAP(L, mu, order, eps, nil)
		check(err)
	case "RLS":
		testName = fmt.Sprintf("%v_%v_L-%v", adfName, applicationName, L)
		af, err = adf.NewFiltRLS(L, mu, eps, nil)
		check(err)
	}

	//d, x := research.MakeXWhiteDData(data, L)

	//d, y, e := run(data, af, L)

	n := len(data)
	var x = make([][]float64, n)
	for i := 0; i < n; i++ {
		x[i] = make([]float64, L)
	}

	d := make([]float64, n)
	y := make([]float64, n)
	e := make([]float64, n)
	var xRow = make([]float64, L)

	for i := 0; i < n; i++ {
		fmt.Printf("working... %d%%\r", (i+1)*100/n)

		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, float64(rand.Intn(math.MaxUint16)-(math.MaxInt16+1)))
		copy(x[i], xRow)
		d[i] = float64(data[i])
		af.Adapt(d[i], x[i])
		y[i] = af.Predict(x[i])
		e[i] = d[i] - y[i]
	}
	//y, e, _, err := af.Run(d, x)

	fmt.Printf("\nwriting to csv...\n")
	check(err)
	research.SaveFilterdDataAsCSV(d, y, e, dataDir, testName)
}

func run(data []int, af adf.AdaptiveFilter, L int) (d, y, e []float64) {
	n := len(data)
	var x = make([][]float64, n)
	for i := 0; i < n; i++ {
		x[i] = make([]float64, L)
	}

	d = make([]float64, n)
	y = make([]float64, n)
	e = make([]float64, n)
	var xRow = make([]float64, L)

	for i := 0; i < n; i++ {
		fmt.Printf("working... %d%%\r", (i+1)*100/n)

		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, float64(rand.Intn(math.MaxUint16)-(math.MaxInt16+1)))
		copy(x[i], xRow)
		d[i] = float64(data[i])
		af.Adapt(d[i], x[i])
		y[i] = af.Predict(x[i])
		e[i] = d[i] - y[i]
	}
	return d, y, e
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
