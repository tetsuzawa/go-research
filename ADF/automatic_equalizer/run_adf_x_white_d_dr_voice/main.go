package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tetsuzawa/go-adflib/adf"
	research "github.com/tetsuzawa/go-research/ADF/raw_drone_convergence"
)

const (
	eps = 1e-5
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

	applicationName := "auto"

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

	fmt.Println("making d, x ...")
	d, x := research.MakeXWhiteDData(data, L)

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
