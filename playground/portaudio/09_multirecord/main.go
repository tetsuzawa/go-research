package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/takuyaohashi/go-wav"
)

//numOutputChannels int, sampleRate float64, framesPerBuffer
const (
	NumOutputChannels = 0
	SampleRate        = 48000
	FramesPerBuffer   = 1024
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	run()
}

var w1 *wav.Writer

func run() {
	args := os.Args
	fmt.Printf( "Record Seconds: %v [sec]\n",args[1])
	RecordSeconds, err := strconv.Atoi(args[1])
	check(err)
	fmt.Printf("Record on %d ch")
	NumInputChannels, err := strconv.Atoi(args[2])
	check(err)

	f1, err := os.Create(`input.wav`)
	check(err)
	defer f1.Close()

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
	paParam.Input.Channels = NumInputChannels
	//paParam.Output.Channels = NumOutputChannels
	paParam.FramesPerBuffer = FramesPerBuffer

	fmt.Println("paParam.Input.Channels", paParam.Input.Channels)
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
		NumChannels:   uint16(NumInputChannels),
		AudioFormat:   1,
	}
	w1, err = wav.NewWriter(f1, p)
	check(err)
	defer w1.Close()

	fmt.Println("recording...")
	err = stream.Start()
	check(err)
	time.Sleep(time.Duration(RecordSeconds) * time.Second)
	err = stream.Stop()
	check(err)
	fmt.Println("end!!")
}

//func callback(inBuf, outBuf []int16) {
func callback(inBuf, outBuf []int16) {
	w1.WriteSamples(inBuf)
}
