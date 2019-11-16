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
	"context"
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
	RecordSeconds, err := strconv.Atoi(args[1])
	fmt.Printf("Record Seconds: %.1f [sec]\n", float64(RecordSeconds))
	check(err)
	NumInputChannels, err := strconv.Atoi(args[2])
	fmt.Printf("Record on %d ch", NumInputChannels)
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

	//////////////////////////////////
	h, err := portaudio.DefaultHostApi()
	check(err)
	paParam := portaudio.LowLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)
	paParam.Input.Channels = NumInputChannels
	//paParam.Output.Channels = NumOutputChannels
	paParam.FramesPerBuffer = FramesPerBuffer * NumInputChannels

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

	inBuf := make([]int16, NumInputChannels*FramesPerBuffer)
	outBuf := make([]int16, NumOutputChannels*FramesPerBuffer)

	//open strea
	stream, err := portaudio.OpenStream(paParam, inBuf, outBuf)
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

	bufCh := make(chan []int16)
	//pass the channel to the main processing function
	ctx, cancel := context.WithCancel(context.Background())

	go wavWrite(ctx, bufCh)

	fmt.Println("recording start")
	err = stream.Start()
	check(err)
	iter := RecordSeconds * SampleRate / FramesPerBuffer
	start := time.Now()
	//var progressRate float64
	for i := 1; i <= iter; i++ {
		/////// progress ///////
		//progressRate = float64(i+1)/float64(iter)
		//fmt.Printf("recording %3d%%, %3.1g [sec]\r", int(float64(i)/float64(iter)*100), float64(i)*float64(RecordSeconds)/float64(iter))
		//fmt.Printf("recording... %3d%%\r", int(float64(i)/float64(iter)*100))
		fmt.Printf("%.1f[sec] : %.1f[sec]\r", time.Since(start).Seconds(), float64(RecordSeconds))
		/////// progress ///////
		err = stream.Read()
		check(err)
		bufCh <- inBuf
		//w1.WriteSamples(inBuf)
	}
	fmt.Printf("\nrecording end\n")
	cancel()

	err = stream.Stop()
	check(err)
	fmt.Printf("\nSuccessfully recorded!!\n")
	fmt.Printf("File saved as `%v`\n", fileName)
}

func wavWrite(ctx context.Context, bufCh chan []int16) {
	for {
		select {
		case buf := <-bufCh:
			w1.WriteSamples(buf)
			continue
		case <-ctx.Done():
			return
		default:
		}
	}
}

//func callback(inBuf, outBuf []int16) {
//	fmt.Println(len(inBuf))
//	w1.WriteSamples(inBuf)
//}
