package main

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
)

func main() {
	err := portaudio.Initialize()
	if err != nil {
		fmt.Println(err)
		err = portaudio.Terminate()
		fmt.Println(err)
		check(err)
	}
	defer portaudio.Terminate()

	devs, err := portaudio.Devices()
	check(err)
	for i, Device := range devs {
		fmt.Printf("\n\n%d\n\n", i)

		fmt.Println("Device.Name", Device.Name)
		fmt.Println("Device.MaxInputChannels", Device.MaxInputChannels)
		fmt.Println("Device.DefaultSampleRate", Device.DefaultSampleRate)
		fmt.Println("Device.HostApi", Device.HostApi)
		fmt.Println("Device.DefaultLowInputLatency", Device.DefaultLowInputLatency)
		fmt.Println("Device.DefaultHighInputLatency", Device.DefaultHighInputLatency)

		fmt.Println("Device.Name", Device.Name)
		fmt.Println("Device.MaxInputChannels", Device.MaxInputChannels)
		fmt.Println("Device.DefaultSampleRate", Device.DefaultSampleRate)
		fmt.Println("Device.HostApi", Device.HostApi)
		fmt.Println("Device.DefaultLowInputLatency", Device.DefaultLowInputLatency)
		fmt.Println("Device.DefaultHighInputLatency", Device.DefaultHighInputLatency)
	}

	//inputDevice, err := portaudio.DefaultInputDevice()
	//check(err)
	//outputDevice, err := portaudio.DefaultOutputDevice()
	//check(err)
}
