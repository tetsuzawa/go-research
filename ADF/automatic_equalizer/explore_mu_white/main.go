package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tetsuzawa/go-adflib/adf"
	research "github.com/tetsuzawa/go-research/ADF/raw_drone_convergence"
)

const (
	eps = 1e-5
)

func main() {

	var (
		L       int
		order   int
		muStart float64
		muEnd   float64
		step    int
	)

	flag.IntVar(&L, "l", 64, "filter length")
	flag.IntVar(&order, "order", 8, "order (AP)")
	flag.Float64Var(&muStart, "start", 1e-5, "start value of mu")
	flag.Float64Var(&muEnd, "end", 2, "end value of mu")
	flag.IntVar(&step, "step", 100, "explore steps")

	flag.Parse()

	fmt.Println("L:", L)
	fmt.Println("order:", order)
	fmt.Println("muStart:", muStart)
	fmt.Println("muEnd:", muEnd)
	fmt.Println("step:", step)

	wavName := flag.Arg(0)
	fmt.Println("wavName:", wavName)
	adfName := flag.Arg(1)
	fmt.Println("adfName:", adfName)
	dataDir := flag.Arg(2)
	fmt.Println("dataDir:", dataDir)

	mu := 1.0
	data := research.ReadDataFromWav(wavName)

	var err error
	var af adf.AdaptiveFilter
	switch adfName {
	case "LMS":
		af, err = adf.NewFiltLMS(L, mu, nil)
		check(err)
	case "NLMS":
		af, err = adf.NewFiltNLMS(L, mu, eps, nil)
		check(err)
	case "AP":
		af, err = adf.NewFiltAP(L, mu, order, eps, nil)
		check(err)
	case "RLS":
		af, err = adf.NewFiltRLS(L, mu, eps, nil)
		check(err)
	}

	fmt.Println("making d, x ...")
	d, x := research.MakeData(data, L)
	fmt.Println("exploring mu ...")
	//mu = ExploreLearning(d, x, af, testName, dataDir)
	mu = research.ExploreLearning(d, x, af, muStart, muEnd, step)
	af.SetStepSize(mu)

	var testName string
	switch adfName {
	case "LMS":
		//testName = fmt.Sprintf("%v_white_mu-%f_L-%v", adfName, mu, L)
		testName = fmt.Sprintf("%v_white_L-%v", adfName, L)
	case "NLMS":
		testName = fmt.Sprintf("%v_white_L-%v", adfName, L)
	case "AP":
		testName = fmt.Sprintf("%v_white_L-%v_order-%v", adfName, L, order)
	case "RLS":
		testName = fmt.Sprintf("%v_white_L-%v", adfName, L)
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
