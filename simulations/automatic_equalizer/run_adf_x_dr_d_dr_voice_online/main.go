package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tetsuzawa/go-adflib/adf"
	"github.com/tetsuzawa/go-adflib/misc"
	simulation "github.com/tetsuzawa/go-research/ADF/automatic_equalizer"
)

const (
	eps = 1e-5
)

func main() {

	jsonName := os.Args[1]
	dataDir := os.Args[2]
	xWavPath := os.Args[3]
	dWavPath := os.Args[4]

	rawJSON, err := ioutil.ReadFile(jsonName)
	check(err)

	var optStepADF = new(simulation.OptStepADF)
	err = json.Unmarshal(rawJSON, optStepADF)
	check(err)

	//wavName := optStepADF.WavName
	adfName := optStepADF.AdfName

	xData := simulation.ReadDataFromWav(xWavPath)
	dData := simulation.ReadDataFromWav(dWavPath)

	L := optStepADF.L
	mu := optStepADF.Mu
	order := optStepADF.Order

	fmt.Println("x wav name:", xWavPath)
	fmt.Println("d wav name:", dWavPath)
	fmt.Println("adf name:", adfName)
	fmt.Println("L:", L)
	fmt.Println("mu:", mu)
	if adfName == "AP" {
		fmt.Println("order:", order)
	}

	var af adf.AdaptiveFilter
	var testName string

	//applicationName = "auto_on_ref_convo"
	applicationName := os.Getenv("BIN_NAME")
	fmt.Printf("\n\napplication name: %v\n\n", applicationName)

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

	//d, x := simulation.MakeXWhiteDData(xData, L)

	//d, y, e := run(xData, af, L)

	n := len(xData)
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
		xRow = append(xRow, float64(xData[i]))
		copy(x[i], xRow)
		d[i] = float64(dData[i])
		af.Adapt(d[i], x[i])
		y[i] = af.Predict(x[i])
		e[i] = d[i] - y[i]
	}
	_, _, w := af.GetParams()

	fmt.Printf("writing to csv...\n")
	simulation.SaveFilterdDataAsCSV(d, y, e, dataDir, testName)
	simulation.SaveWAsCSV(w, dataDir, testName)
	fmt.Printf("writing to wav...\n")
	simulation.SaveFilteredDataToWav(e, dataDir, testName)
	fmt.Printf("\n")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
