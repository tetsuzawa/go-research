package main

import (
	"encoding/binary"
	"gonum.org/v1/gonum/floats"
	"log"
	"math"
)
func LoggingSettings(logFile string)  {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}


func float32ToBytes(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return b
}

func float32sToBytes(fs []float32) []byte {
	bs := make([]byte, len(fs)*4)
	b := make([]byte, 4)
	for _, f := range fs {
		bits := math.Float32bits(f)
		binary.LittleEndian.PutUint32(b, bits)
		bs = append(bs, b...)
	}
	return bs

}

func Normalize(fs []float32) []float32 {
	fs64 := make([]float64, len(fs))
	for i, s := range fs {
		fs64[i] = float64(s)
	}
	m := floats.Max(fs64)
	for i, s := range fs64 {
		fs[i] = float32(s / m)
	}
	return fs
}

func Float32sToInt16s(fs []float32) []int16 {
	//fs = Normalize(fs)
	is := make([]int16, len(fs))
	for i, s := range fs {
		is[i] = int16(s * math.MaxInt16)
	}
	return is
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func Float64sToInt16s(fs []float64) []int16 {
	is := make([]int16, len(fs))
	for i, s := range fs {
		is[i] = int16(s * math.MaxInt16)
	}
	return is
}

func Int16sToFloat64s(is []int16) []float64 {
	//is = NormalizeInt16s(is)
	fs := make([]float64, len(is))
	for i, s := range is {
		//fs[i] = float64(s) / math.MaxInt16 * math.MaxFloat64
		fs[i] = float64(s)
	}
	return fs
}
