package main

import (
	"machine"
	"time"
)

func main() {
	// On board LED is connected to GPIO 2
	led := machine.Pin(2)
	// Configure PIN as output
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Infinite main loop
	for {
		// Turn LED off
		led.Low()
		// Wait for 1 second
		time.Sleep(time.Millisecond * 1000)
		// Turn LED on
		led.High()
		// Wait for 1 second
		time.Sleep(time.Millisecond * 1000)
	}
}

