package main

import (
	"encoding/json"
	"fmt"
	"github.com/tetsuzawa/go-adflib/adf"
	research "github.com/tetsuzawa/go-research/ADF/raw_drone_convergence"
	"io/ioutil"
	"os"
)

const (
	eps = 1e-5
	w   = "zeros"
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

	var af adf.AdaptiveFilter
	var testName string
	switch adfName {
	case "LMS":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v", adfName, mu, L)
		af, err = adf.NewFiltLMS(L, mu, w)
		check(err)
	case "NLMS":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v", adfName, mu, L)
		af, err = adf.NewFiltNLMS(L, mu, eps, w)
		check(err)
	case "AP":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v_order-%v", adfName, mu, L, order)
		af, err = adf.NewFiltAP(L, mu, order, eps, w)
		check(err)
	case "RLS":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v", adfName, mu, L)
		af, err = adf.NewFiltRLS(L, mu, eps, w)
		check(err)
	}

	fmt.Println("making d, x ...")
	d, x := research.MakeData(data, L)

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
