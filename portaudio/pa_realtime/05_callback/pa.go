package main

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"github.com/takuyaohashi/go-wav"
	"log"
	"os"
	"time"
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

var w *wav.Writer

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

	fmt.Println("paParam.SampleRate", paParam.SampleRate)
	fmt.Println("paParam.FramesPerBuffer", paParam.FramesPerBuffer)

	fmt.Println("paParam.Input.Device.Name", paParam.Input.Device.Name)
	fmt.Println("paParam.Input.Device.MaxInputChannels", paParam.Input.Device.MaxInputChannels)
	fmt.Println("paParam.Input.Device.DefaultSampleRate", paParam.Input.Device.DefaultSampleRate)
	fmt.Println("paParam.Input.Device.HostApi", paParam.Input.Device.HostApi)
	fmt.Println("paParam.Input.Device.DefaultLowInputLatency", paParam.Input.Device.DefaultLowInputLatency)
	fmt.Println("paParam.Input.Device.DefaultHighInputLatency", paParam.Input.Device.DefaultHighInputLatency)

	fmt.Println("paParam.Output.Device.Name", paParam.Output.Device.Name)
	fmt.Println("paParam.Output.Device.MaxInputChannels", paParam.Output.Device.MaxInputChannels)
	fmt.Println("paParam.Output.Device.DefaultSampleRate", paParam.Output.Device.DefaultSampleRate)
	fmt.Println("paParam.Output.Device.HostApi", paParam.Output.Device.HostApi)
	fmt.Println("paParam.Output.Device.DefaultLowInputLatency", paParam.Output.Device.DefaultLowInputLatency)
	fmt.Println("paParam.Output.Device.DefaultHighInputLatency", paParam.Output.Device.DefaultHighInputLatency)


	//open stream
	stream, err := portaudio.OpenStream(paParam, callback)
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
	w, err = wav.NewWriter(f, p)
	defer w.Close()

	fmt.Println("recording...")
	err = stream.Start()
	check(err)
	time.Sleep(5 * time.Second)
	err = stream.Stop()
	check(err)
	fmt.Println("end!!")
}

func callback(inBuf, outBuf []int16) {
	for i, _ := range outBuf {
		outBuf[i] = inBuf[i]
	}
	//copy(outBuf, inBuf)
	w.WriteSamples(inBuf)
}
