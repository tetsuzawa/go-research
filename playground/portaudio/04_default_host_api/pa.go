package main

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"github.com/takuyaohashi/go-wav"
	"log"
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

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	run()
}

func run() {
	f, err := os.Create(`input.wav`)
	check(err)
	defer f.Close()

	err = portaudio.Initialize()
	if err != nil {
		fmt.Println(err)
		err = portaudio.Terminate()
		fmt.Println(err)
		check(err)
	}
	defer portaudio.Terminate()

	//////////////////////////////////
	h, err := portaudio.DefaultHostApi()
	check(err)
	paParam := portaudio.LowLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)
	paParam.Input.Channels = 1
	paParam.Output.Channels = 1

	fmt.Println(paParam.Input.Device.Name)
	fmt.Println(paParam.Input.Device.MaxInputChannels)
	fmt.Println(paParam.Input.Device.DefaultSampleRate)
	fmt.Println(paParam.Input.Device.HostApi)
	fmt.Println(paParam.Input.Device.DefaultLowInputLatency)
	fmt.Println(paParam.Input.Device.DefaultHighInputLatency)

	fmt.Println(paParam.Output.Device.Name)
	fmt.Println(paParam.Output.Device.MaxOutputChannels)
	fmt.Println(paParam.Output.Device.DefaultSampleRate)
	fmt.Println(paParam.Output.Device.HostApi)
	fmt.Println(paParam.Output.Device.DefaultLowInputLatency)
	fmt.Println(paParam.Output.Device.DefaultHighInputLatency)

	bufferIn := make([]int16, FramesPerBuffer)
	bufferOut := make([]int16, FramesPerBuffer)

	//open stream
	stream, err := portaudio.OpenStream(paParam, bufferIn, bufferOut)
	check(err)
	defer stream.Close()
	fmt.Println("info: ", stream.Info())

	//make input.wav
	p := wav.WriterParam{
		SampleRate:    48000,
		BitsPerSample: 16,
		NumChannels:   1,
		AudioFormat:   1,
	}
	w, err := wav.NewWriter(f, p)
	defer w.Close()

	fmt.Println("recording...")
	err = stream.Start()
	check(err)
	//iter := int(float64(5) / (float64(FramesPerBuffer) / float64(SampleRate)))
	iter := 235
	for i := 0; i < iter; i++ {
		err = stream.Read()
		check(err)
		copy(bufferOut, bufferIn)
		//for i, _ := range bufferOut {
		//	bufferOut[i] = bufferIn[i]
		//}
		err = stream.Write()
		check(err)
		w.WriteSamples(bufferIn)
	}
	//fmt.Println(bufferIn)
	fmt.Println("end!!")
	//err = stream.Write()
	//if err != nil {
	//	fmt.Println(err)
	//}
	err = stream.Stop()
	check(err)

}
