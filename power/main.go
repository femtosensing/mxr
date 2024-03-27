package main

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
)

var POWER_EN_PIN = 26

func PowerOn() error {
	err := rpio.Open()
	if err != nil {
		return err
	}

	PowerPin := rpio.Pin(POWER_EN_PIN)
	PowerPin.Output()
	PowerPin.High()
	return nil
}

func PowerOff() error {
	err := rpio.Open()
	if err != nil {
		return err
	}

	PowerPin := rpio.Pin(POWER_EN_PIN)
	PowerPin.Output()
	PowerPin.Low()

	return nil
}

func main() {
	args := os.Args
	on := true
	if len(args) > 1 {
		if args[1] == "-off" {
			on = false
		}
	}

	var err error
	if on {
		err = PowerOn()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Power On")
	} else {
		err = PowerOff()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Power Off")
	}
}
