package _1_allbuf

import (
	"encoding/binary"
	"fmt"
	"github.com/gordonklaus/portaudio"
	"github.com/takuyaohashi/go-wav"
	"gonum.org/v1/gonum/floats"
	"log"
	"math"
	"os"
)

//numOutputChannels int, sampleRate float64, framesPerBuffer
const (
	RecordSeconds int = 5

	NumInputChannels  int = 1
	NumOutputChannels int = 0
	SampleRate        int = 48000
	FramesPerBuffer   int = 1024
)

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
	fs = Normalize(fs)
	is := make([]int16, len(fs))
	for i, s := range fs {
		is[i] = int16(s * math.MaxInt16)
	}
	return is
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	f, err := os.Create(`input.wav`)
	check(err)
	defer f.Close()

	err = portaudio.Initialize()
	check(err)
	defer portaudio.Terminate()
	inputDevice, err := portaudio.DefaultInputDevice()
	check(err)
	outputDevice, err := portaudio.DefaultOutputDevice()
	check(err)

	streamParameters := portaudio.StreamParameters{
		Input: portaudio.StreamDeviceParameters{inputDevice, 1, inputDevice.DefaultLowInputLatency},
		//Output: portaudio.StreamDeviceParameters{outputDevice, 1, outputDevice.DefaultHighOutputLatency},
		Output:          portaudio.StreamDeviceParameters{nil, 0, outputDevice.DefaultLowOutputLatency},
		SampleRate:      48000,
		FramesPerBuffer: 1024,
		Flags:           portaudio.NoFlag,
	}
	fmt.Println(streamParameters.Input.Device.Name)
	fmt.Println(streamParameters.Input.Device.MaxInputChannels)
	fmt.Println(streamParameters.Input.Device.DefaultSampleRate)
	fmt.Println(streamParameters.Input.Device.HostApi)
	fmt.Println(streamParameters.Input.Device.DefaultLowInputLatency)
	fmt.Println(streamParameters.Input.Device.DefaultHighInputLatency)

	//fmt.Println(streamParameters.Output.Device.Name)
	//fmt.Println(streamParameters.Output.Device.MaxOutputChannels)
	//fmt.Println(streamParameters.Output.Device.DefaultSampleRate)
	//fmt.Println(streamParameters.Output.Device.HostApi)
	//fmt.Println(streamParameters.Output.Device.DefaultLowInputLatency)
	//fmt.Println(streamParameters.Output.Device.DefaultHighInputLatency)

	bufferIn := make([]float32, SampleRate*RecordSeconds)
	bufferOut := make([]float32, SampleRate*RecordSeconds)
	//bufferIn := make([]float32, FramesPerBuffer)
	//bufferOut := make([]float32, FramesPerBuffer)

	//stream, err := portaudio.OpenDefaultStream(
	//	NumInputChannels,
	//	NumOutputChannels,
	//	float64(SampleRate),
	//	len(buffer),
	//	func(in []float32) {
	//		for i, _ := range buffer {
	//			buffer[i] = in[i]
	//		}
	//	})
	//stream, err := portaudio.OpenDefaultStream(
	//	NumInputChannels,
	//	NumOutputChannels,
	//	float64(SampleRate),
	//	FramesPerBuffer,
	//	)
	stream, err := portaudio.OpenStream(streamParameters, bufferIn, bufferOut)
	check(err)
	defer stream.Close()

	fmt.Println("recording...")
	err = stream.Start()
	check(err)
	err = stream.Read()
	check(err)
	//fmt.Println(bufferIn)
	fmt.Println("end!!")
	//err = stream.Write()
	//if err != nil {
	//	fmt.Println(err)
	//}
	err = stream.Stop()
	check(err)

	//fmt.Println("recording...")
	//for i := 0; i < SampleRate*RecordSeconds/FramesPerBuffer; i++ {
	//	//err = stream.Write()
	//	//if err != nil {
	//	//	panic(err)
	//		//fmt.Println(err)
	//	//}
	//	err = stream.Read()
	//	if err != nil {
	//		//panic(err)
	//		fmt.Println(err)
	//	}
	//}
	//fmt.Println("end!!")
	//err = stream.Stop()
	//if err != nil {
	//	panic(err)
	//}
	//}
	//tmpBuff := bufferIn
	//func() {
	//	by := float32sToBytes(tmpBuff)
	//	fmt.Println(by)
	//	i, err = f.Write(by)
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
	//err = stream.Write()
	//if err != nil {
	//
	//	log.Println(err)
	//}
	p := wav.WriterParam{
		SampleRate:    48000,
		BitsPerSample: 16,
		NumChannels:   1,
		AudioFormat:   1,
	}
	w, err := wav.NewWriter(f, p)
	defer w.Close()
	is := Float32sToInt16s(bufferIn)
	w.WriteSamples(is)
}
