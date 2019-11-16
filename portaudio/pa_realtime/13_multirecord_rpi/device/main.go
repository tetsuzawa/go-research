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
		fmt.Printf("\n\n%d\n", i)

		fmt.Printf("\n\nInput params\n\n")

		fmt.Println("Device.Name", Device.Name)
		fmt.Println("Device.MaxInputChannels", Device.MaxInputChannels)
		fmt.Println("Device.DefaultSampleRate", Device.DefaultSampleRate)
		fmt.Println("Device.HostApi", Device.HostApi)
		fmt.Println("Device.DefaultLowInputLatency", Device.DefaultLowInputLatency)
		fmt.Println("Device.DefaultHighInputLatency", Device.DefaultHighInputLatency)

		fmt.Printf("\n\nOutput params\n\n")

		fmt.Println("Device.Name", Device.Name)
		fmt.Println("Device.MaxOutputChannels", Device.MaxOutputChannels)
		fmt.Println("Device.DefaultSampleRate", Device.DefaultSampleRate)
		fmt.Println("Device.HostApi", Device.HostApi)
		fmt.Println("Device.DefaultLowOutputLatency", Device.DefaultLowOutputLatency)
		fmt.Println("Device.DefaultHighOutputLatency", Device.DefaultHighOutputLatency)
	}

	//inputDevice, err := portaudio.DefaultInputDevice()
	//check(err)
	//outputDevice, err := portaudio.DefaultOutputDevice()
	//check(err)
}
