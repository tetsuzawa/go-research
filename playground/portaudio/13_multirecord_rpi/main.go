/*
Contents: wav to DSB converter.
	This program converts .wav file to .DSB files.
	Please run `wav_to_DSB --help` for details.
Usage: wav_to_DSB (-o /path/to/out.DSB) /path/to/file.wav
Author: Tetsu Takizawa
E-mail: tt15219@tomakomai.kosen-ac.jp
LastUpdate: 2019/11/16
DateCreated  : 2019/11/16
*/
package main

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"github.com/tetsuzawa/go-wav"
	"log"
	"os"
	"strconv"
	"time"
)

//numOutputChannels int, sampleRate float64, framesPerBuffer
const (
	FramesPerBuffer = 1024
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	run()
}

var w1 *wav.Writer

func run() {
	args := os.Args
	RecordSeconds, err := strconv.Atoi(args[1])
	fmt.Printf("Record Seconds: %.1f [sec]\n", float64(RecordSeconds))
	check(err)
	NumInputChannels, err := strconv.Atoi(args[2])
	fmt.Printf("Record on %d ch", NumInputChannels)
	check(err)

	NumOutputChannels, err := strconv.Atoi(args[3])
	fmt.Printf("Play on %d ch", NumOutputChannels)
	check(err)

	SampleRate, err := strconv.Atoi(args[4])
	fmt.Printf("Recording in sample rate %d", SampleRate)
	check(err)

	fileName := `input.wav`
	f1, err := os.Create(fileName)
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

	devices, err := portaudio.Devices()
	check(err)
	//
	inputDevice := devices[0]
	//inputDevice, err := portaudio.DefaultInputDevice()
	//check(err)
	//outputDevice := devices[2]
	outputDevice, err := portaudio.DefaultOutputDevice()
	check(err)

	fmt.Println(inputDevice.Name, outputDevice.Name)
	//if inputDevice.Name != "snd_rpi_pcm1808_adc: - (hw:0,0)" || outputDevice.Name !=  "bcm2835 ALSA: IEC958/HDMI (hw:1,1)"{
	//	fmt.Println("Device incorrect")
	//	os.Exit(1)
	//}

	paParam := portaudio.StreamParameters{
		Input:  portaudio.StreamDeviceParameters{inputDevice, NumInputChannels, inputDevice.DefaultLowInputLatency},
		Output: portaudio.StreamDeviceParameters{outputDevice, NumOutputChannels, outputDevice.DefaultLowOutputLatency},
		//Output:          portaudio.StreamDeviceParameters{nil, 0, outputDevice.DefaultLowOutputLatency},
		SampleRate:      float64(SampleRate),
		FramesPerBuffer: FramesPerBuffer * NumInputChannels,
		Flags:           portaudio.NoFlag,
	}

	//////////////////////////////////
	//h, err := portaudio.DefaultHostApi()
	//check(err)
	//paParam := portaudio.LowLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)
	//paParam.Input.Channels = NumInputChannels
	//paParam.Output.Channels = NumOutputChannels
	//paParam.FramesPerBuffer = FramesPerBuffer * NumInputChannels

	fmt.Println("paParam.Input.Channels", paParam.Input.Channels)
	fmt.Println("paParam.Output.Channels", paParam.Output.Channels)
	fmt.Println("paParam.SampleRate", paParam.SampleRate)
	fmt.Println("paParam.FramesPerBuffer", paParam.FramesPerBuffer)

	fmt.Printf("\n\nInput params\n\n")

	fmt.Println("paParam.Input.Device.Name", paParam.Input.Device.Name)
	fmt.Println("paParam.Input.Device.MaxInputChannels", paParam.Input.Device.MaxInputChannels)
	fmt.Println("paParam.Input.Device.DefaultSampleRate", paParam.Input.Device.DefaultSampleRate)
	fmt.Println("paParam.Input.Device.HostApi", paParam.Input.Device.HostApi)
	fmt.Println("paParam.Input.Device.DefaultLowInputLatency", paParam.Input.Device.DefaultLowInputLatency)
	fmt.Println("paParam.Input.Device.DefaultHighInputLatency", paParam.Input.Device.DefaultHighInputLatency)

	fmt.Printf("\n\nOutput params\n\n")

	fmt.Println("paParam.Output.Device.Name", paParam.Output.Device.Name)
	fmt.Println("paParam.Output.Device.MaxOutputChannels", paParam.Output.Device.MaxOutputChannels)
	fmt.Println("paParam.Output.Device.DefaultSampleRate", paParam.Output.Device.DefaultSampleRate)
	fmt.Println("paParam.Output.Device.HostApi", paParam.Output.Device.HostApi)
	fmt.Println("paParam.Output.Device.DefaultLowOutputLatency", paParam.Output.Device.DefaultLowOutputLatency)
	fmt.Println("paParam.Output.Device.DefaultHighOutputLatency", paParam.Output.Device.DefaultHighOutputLatency)

	inBuf := make([]int16, NumInputChannels*FramesPerBuffer)
	outBuf := make([]int16, NumOutputChannels*FramesPerBuffer)

	//open strea
	stream, err := portaudio.OpenStream(paParam, inBuf, outBuf)
	check(err)
	defer stream.Close()
	fmt.Println("info: ", stream.Info())

	//make input.wav
	p := wav.WriterParam{
		SampleRate:    uint32(SampleRate),
		BitsPerSample: 16,
		NumChannels:   uint16(NumInputChannels),
		AudioFormat:   1,
	}
	w1, err = wav.NewWriter(f1, p)
	check(err)
	defer w1.Close()

	fmt.Println("recording start")
	err = stream.Start()
	check(err)
	start := time.Now()
	iter := RecordSeconds * SampleRate / FramesPerBuffer
	for i := 1; i <= iter; i++ {
		/////// progress ///////
		fmt.Printf("%.1f[sec] : %.1f[sec]\r", time.Since(start).Seconds(), float64(RecordSeconds))
		err = stream.Read()
		check(err)
		w1.WriteSamples(inBuf)
	}

	fmt.Printf("\nrecording end\n")

	err = stream.Stop()
	check(err)
	fmt.Printf("\nSuccessfully recorded!!\n")
	fmt.Printf("File saved as `%v`\n", fileName)
}
