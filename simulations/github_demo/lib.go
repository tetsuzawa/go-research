package research

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-audio/wav"
	"github.com/tetsuzawa/go-adflib/adf"
	"github.com/tetsuzawa/go-adflib/misc"
	"gonum.org/v1/gonum/floats"
)

type OptStepADF struct {
	WavName string  `json:"wav_name"`
	AdfName string  `json:"adf_name"`
	Mu      float64 `json:"mu"`
	L       int     `json:"l"`
	Order   int     `json:"order"`
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

func MakeData(data []int, L int) (d []float64, x [][]float64) {
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

func ExploreLearning(d []float64, x [][]float64, af adf.AdaptiveFilter, muStart, muEnd float64, steps int) float64 {

	//es, mus, err := adf.ExploreLearning(af, d, x, 0.00001, 2.0, 100, 0.5, 1, "MSE", nil)
	es, mus, err := adf.ExploreLearning(af, d, x, muStart, muEnd, steps, 0.5, 1, "MSE", nil)
	check(err)

	res := make(map[float64]float64, len(es))
	for i := 0; i < len(es); i++ {
		res[es[i]] = mus[i]
	}
	eMin := floats.Min(es)
	fmt.Printf("the step size mu with the smallest error is %.3f\n", res[eMin])

	//fw, err := os.Create(filepath.Join(dataDir, testName+"_opt_mu.log"))
	//check(err)
	//_, err = fmt.Fprintf(fw, "%g\n", res[eMin])
	//check(err)
	//err = fw.Close()
	//check(err)

	return res[eMin]
}

func SaveFilterdDataAsCSV(d, y, e []float64, dataDir string, testName string) {
	n := len(d)
	fw, err := os.Create(filepath.Join(dataDir, testName+".csv"))
	check(err)
	writer := bufio.NewWriter(fw)
	for i := 0; i < n; i++ {
		//_, err = fw.Write([]byte(fmt.Sprintf("%f,%f,%f\n", d[i], y[i], e[i])))
		//_, err = writer.WriteString(fmt.Sprintf("%g,%g,%g\n", d[i], y[i], e[i]))
		_, err = fmt.Fprintf(writer, "%g,%g,%g\n", d[i], y[i], e[i])
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
