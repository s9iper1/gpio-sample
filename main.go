package main

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	pin := rpio.Pin(21)
	pin.Output()

	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}

}
