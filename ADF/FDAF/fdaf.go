package main

import (
	"encoding/binary"
	"fmt"
	"github.com/mjibson/go-dsp/fft"
	"github.com/takuyaohashi/go-wav"
	"github.com/tetsuzawa/converter"
	"github.com/tetsuzawa/go-3daudio/utils"
	"log"
	"math/cmplx"
	"os"
	"path/filepath"
	"reflect"
)

const (
	mu = 0.001
	L  = 128
)

func ExtractFileName(path string) string {
	// Fixed with a nice method given by mattn-san
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func unset(s []float64, i int) []float64 {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func fdaf(data []float64) []float64 {
	// 1 w(o). random value. use as vector.
	var w = make([]float64, 2*L)
	// output buffer
	var err_buf = make([]float64, 0)

	var u = make([]float64, 2*L)
	var zeros = make([]float64, L)

	var idx int
	Y := make([]complex128, 2*L)
	y := make([]float64, L)
	e := make([]float64, L)
	EU := make([]complex128, 2*L)

FDAF:
	for {
		fmt.Printf("Filter adapting... %d%% \r", (idx+1)*100/len(data))

		// 2.0 Initialize phi = 0s

		// 2.1 Iterate for i = 0, 1, 2, 3, ..., L-1 (k is the block index)
		for j := 0; j < L; j++ {

			// 2.1.0 Read/generate a new data pair
			in := data[idx]
			if idx == len(data)-1 {
				fmt.Printf("\nAdaptation completed!!\n")
				break FDAF
			}

			u = unset(u, 0)
			u = append(u, in)

			j++
			idx++
		}

		// 1 compute the output of the filter for the block kM, ..., KM + M -1
		W := fft.FFT(converter.Float64sToComplex128s(append(w[:L], zeros...)))
		U := fft.FFT(converter.Float64sToComplex128s(u))
		for i := 0; i < 2*L; i++ {
			Y[i] = W[i] * U[i]
		}
		y_raw := fft.IFFT(Y)[L:]
		for i := 0; i < L; i++ {
			y[i] = real(y_raw[i])
			e[i] = u[i] - y[i]
		}

		// 2 compute the correlation vector
		aux1 := fft.FFT(converter.Float64sToComplex128s(append(zeros, e...)))
		aux2 := fft.FFT(converter.Float64sToComplex128s(u))
		for i := 0; i < 2*L; i++ {
			EU[i] = aux1[i] * cmplx.Conj(aux2[i])
		}
		phi := fft.IFFT(EU)[:L]

		// 3 update the parameters of the filter
		aux1 = fft.FFT(converter.Float64sToComplex128s(append(w[:L], zeros...)))
		aux2 = fft.FFT(append(phi, converter.Float64sToComplex128s(zeros)...))
		for i := 0; i < 2*L; i++ {
			W[i] = aux1[i] + mu*aux2[i]
		}
		aux3 := fft.IFFT(W)
		for i := 0; i < 2*L; i++ {
			w[i] = real(aux3[i])
		}
		err_buf = append(err_buf, e...)
	}

	return err_buf
}

func main() {
	utils.LoggingSettings("fdaf.log")

	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	w, err := wav.NewReader(f)
	if err != nil {
		log.Fatalln(err)
	}

	var data interface{}
	data, err = w.ReadSamples(int(w.GetSubChunkSize()) / int(w.GetNumChannels()) / int(w.GetBitsPerSample()) * 8)
	if err != nil {
		log.Fatalln(err)
	}

	if reflect.TypeOf(data) != reflect.TypeOf([]int16{0,}) {
		log.Fatalln(err)
	}

	value, ok := data.([]int16)
	if !ok {
		log.Fatalln("Data type is not valid")
	}

	estErr := fdaf(converter.Int16sToFloat64s(value))

	inFileName := ExtractFileName(fileName)
	wav_out_dir := "/Users/tetsu/personal_files/Research/filters/test/FDAF_wav/"
	wav_out_name := fmt.Sprintf("%s_mu-%f_L-%d.wav", inFileName, mu, L)

	fw, err := os.Create(wav_out_dir + wav_out_name)
	if err != nil {
		log.Fatalln(err)
	}
	defer fw.Close()

	p := wav.WriterParam{
		SampleRate:    48000,
		BitsPerSample: 16,
		NumChannels:   1,
		AudioFormat:   1,
	}
	ww, err := wav.NewWriter(fw, p)
	if err != nil {
		log.Fatalln(err)
	}
	ww.WriteSamples(converter.Float64sToInt16s(estErr))

	b := make([]byte, 2)
	buf := make([]byte, 0)

	for i, v := range value {
		fmt.Printf("Writing data to buffer... %d%%\r", (i+1)*100/len(value))

		ui := converter.Int16ToUint16(v)
		binary.LittleEndian.PutUint16(b, ui)
		buf = append(buf, b...)
	}
	fmt.Printf("\nWriting completed!!\n")

	_, err = fw.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n\n")
	fmt.Println("end!!")
}
