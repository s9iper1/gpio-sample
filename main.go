package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	// Use mcu pin 10, corresponds to physical pin 21 on the pi
	pin = rpio.Pin(21)
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// for my own understand defer is for executing the item after all functions have been excuted. so we call close
	// when code is exucuted
	defer rpio.Close()

	// Set pin to output mode as it will be a out call
	pin.Output()

	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}

}
