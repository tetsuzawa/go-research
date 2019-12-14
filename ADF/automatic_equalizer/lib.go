package research

import (
	"bufio"
	"fmt"
	"github.com/go-audio/wav"
	"github.com/tetsuzawa/go-adflib/adf"
	"github.com/tetsuzawa/go-adflib/misc"
	"gonum.org/v1/gonum/floats"
	"math"
	"math/rand"
	"os"
	"path/filepath"
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
	fmt.Printf("SourceBitDepth: %v\n", buf.SourceBitDepth)

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

func MakeXWhiteDData(data []int, L int) (d []float64, x [][]float64) {
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
		xRow = append(xRow, float64(rand.Intn(math.MaxUint16)-(math.MaxInt16+1)))
		copy(x[i], xRow)
		d[i] = float64(data[i])
	}
	return d, x
}

func MakeDataWithNoise(data []int, L int) (d []float64, x [][]float64) {
	rand.Seed(1)
	n := len(data)
	//input value
	x = make([][]float64, n)
	for i := 0; i < n; i++ {
		x[i] = make([]float64, L)
	}
	//noise

	//read filter
	//fr, err := os.Open("../csvfiles/lp_filt_128.csv")
	//defer fr.Close()
	//check(err)
	//sc := bufio.NewScanner(fr)
	//var f []float64
	//var v float64
	//for sc.Scan() {
	//	v, err = strconv.ParseFloat(sc.Text(), 64)
	//	check(err)
	//	f = append(f, v)
	//}
	tapNum := 4
	f := make([]float64, tapNum)
	f[tapNum-1] = 1.0

	//desired value
	d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = misc.Unset(xRow, 0)
		//fmt.Println("data[i]", i, data[i])
		xRow = append(xRow, float64(data[i]))
		copy(x[i], xRow)
		//d[i] = float64(data[i])+0.05*(float64(rand.Int31()-math.MaxInt32/2)/2)
		d[i] = float64(data[i])
	}
	dc := Convolve(d, f)
	//fmt.Println(dc)

	for i := 0; i < n; i++ {
		d[i] = dc[i] + 0.00316*(float64(rand.Int31())/float64(math.MaxInt16)-math.MaxInt16)
	}

	return d, x
}

func Convolve(xs, ys []float64) []float64 {
	var convLen, sumLen = len(xs), len(ys)
	if convLen > sumLen {
		ys = append(ys, make([]float64, convLen-sumLen)...)
	} else {
		convLen, sumLen = sumLen, convLen
		xs = append(xs, make([]float64, convLen-sumLen)...)
	}
	var rs = make([]float64, convLen)
	//fmt.Println(convLen)
	//fmt.Println(sumLen)
	var nodeSum float64
	var i, j int
	for i = 0; i < convLen; i++ {
		for j = 0; j < sumLen; j++ {
			if i-j < 0 {
				continue
			}
			nodeSum += xs[i-j] * ys[j]
		}
		rs[i] = nodeSum
		nodeSum = 0
	}
	return rs
}

func ExploreLearning(d []float64, x [][]float64, af adf.AdaptiveFilter, muStart, muEnd float64, steps int) float64 {

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
