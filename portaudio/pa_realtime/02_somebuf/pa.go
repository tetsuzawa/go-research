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

	//bufferIn := make([]float32, SampleRate*RecordSeconds)
	//bufferOut := make([]float32, SampleRate*RecordSeconds)
	bufferIn := make([]float32, FramesPerBuffer)
	bufferOut := make([]float32, FramesPerBuffer)

	//open stream
	stream, err := portaudio.OpenStream(streamParameters, bufferIn, bufferOut)
	check(err)
	defer stream.Close()

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
		is := Float32sToInt16s(bufferIn)
		w.WriteSamples(is)
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
