package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/tetsuzawa/go-adflib/adf"
	research "github.com/tetsuzawa/go-research/ADF/raw_drone_convergence"
)

const (
	eps = 1e-5
	w   = "zeros"
)

func main() {

	wavName := os.Args[1]
	adfName := os.Args[2]
	dataDir := os.Args[3]

	data := research.ReadDataFromWav(wavName)
	L, err := strconv.Atoi(os.Args[3])
	check(err)
	mu := 1.0
	order, err := strconv.Atoi(os.Args[4])
	check(err)

	var af adf.AdaptiveFilter
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

	fmt.Println("making d, x ...")
	d, x := research.MakeData(data, L)
	fmt.Println("exploring mu ...")
	//mu = ExploreLearning(d, x, af, testName, dataDir)
	mu = research.ExploreLearning(d, x, af, 0.000001, 2, 100)
	af.SetStepSize(mu)

	var testName string
	switch adfName {
	case "LMS":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v", adfName, mu, L)
	case "NLMS":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v", adfName, mu, L)
	case "AP":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v_order-%v", adfName, mu, L, order)
	case "RLS":
		testName = fmt.Sprintf("%v_static_mu-%f_L-%v", adfName, mu, L)
	}

	var optadf = &research.OptStepADF{
		WavName: wavName,
		AdfName: adfName,
		Mu:      mu,
		L:       L,
		Order:   order,
	}

	outadfJSON, err := json.Marshal(optadf)
	check(err)
	fw, err := os.Create(filepath.Join(dataDir, testName+".json"))
	check(err)
	defer fw.Close()
	_, err = fw.Write(outadfJSON)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
