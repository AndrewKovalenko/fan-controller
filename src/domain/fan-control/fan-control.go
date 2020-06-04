package fanControl

import (
	"errors"

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

func Init() error {
	err := gpio.Open()

	if err != nil {
		return err
	}

	raspiCoolingFan.controlPin.Mode(gpio.Pwm)
	raspiCoolingFan.controlPin.Freq(pmwClockFrequency)
	raspiCoolingFan.controlPin.DutyCycle(fanOff, cycleLength)

	gpio.StartPwm()
	return nil
}

func SetFanSpeed(percentage uint8) error {
	if percentage < 0 || percentage > 100 {
		return errors.New("Fan speed can't be greater than 100 and less than 0")
	}

	raspiCoolingFan.controlPin.DutyCycle(uint32(percentage), cycleLength)
	return nil
}

func CleanUp() {
	gpio.StopPwm()
	gpio.Close()
}
