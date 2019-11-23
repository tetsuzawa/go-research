package main


import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/tetsuzawa/go-adflib/adf"
	research "github.com/tetsuzawa/go-research/ADF/raw_drone_convergence"
)

const (
	eps = 1e-5
	w   = "zeros"
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

	jsonName := os.Args[1]
	rawJSON ,err := ioutil.ReadFile(jsonName)
	check(err)
	


	wavName = os.Args[1]
	data := research.ReadDataFromWav(wavName)

	adfName = os.Args[2]

	L, err = strconv.Atoi(os.Args[3])
	check(err)
	mu, err = strconv.ParseFloat(os.Args[4], 64)
	check(err)
	order, err = strconv.Atoi(os.Args[5])
	check(err)

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

	dataDir = "data"

	fmt.Println("making d, x ...")
	d, x := research.MakeData(data, L)
	//fmt.Println("exploring mu ...")
	//mu = ExploreLearning(d, x, af, testName, dataDir)
	//mu = research.ExploreLearning(d, x, af, 0.000001, 2, 100)
	//af.SetStepSize(mu)

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
