package fanControl

import (
	"log"

	gpio "github.com/stianeikeland/go-rpio/v4"
)

const pmwFanControlPin = 18
const cycleLength = 100
const desiredPmwFrequency = 50 // 50Hz
const pmwClockFrequency = desiredPmwFrequency * cycleLength
const fanOff = 0

type fan struct {
	controlPin gpio.Pin
}

var raspiCoolingFan = fan{controlPin: gpio.Pin(pmwFanControlPin)}

func Init() {
	err := gpio.Open()

	if err != nil {
		log.Fatal("Unable to initialize GPIO pin 18")
	}

	raspiCoolingFan.controlPin = gpio.Pin(pmwFanControlPin)
	raspiCoolingFan.controlPin.Mode(gpio.Pwm)
	raspiCoolingFan.controlPin.Freq(pmwClockFrequency)
	raspiCoolingFan.controlPin.DutyCycle(fanOff, cycleLength)

	gpio.StartPwm()
}

func SetFanSpeed(percentage uint32) {
	if percentage < 0 || percentage > 100 {
		log.Println("Wrong percentage value")
		return
	}

	raspiCoolingFan.controlPin.DutyCycle(percentage, cycleLength)

}

func CleanUp() {
	gpio.StopPwm()
	gpio.Close()
}
