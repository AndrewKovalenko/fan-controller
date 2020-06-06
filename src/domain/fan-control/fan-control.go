package fanControl

import (
	"errors"

	gpio "github.com/stianeikeland/go-rpio/v4"
)

const pmwFanControlPin = 18
const cycleLength = 100
const desiredPmwFrequency = 50 // 5kHz
const pmwClockFrequency = desiredPmwFrequency * cycleLength
const fanOff = 0

var fanControllingPin = gpio.Pin(pmwFanControlPin)

func Init() error {
	if err := gpio.Open(); err != nil {
		return err
	}

	fanControllingPin.Pwm()

	fanControllingPin.Freq(pmwClockFrequency)
	fanControllingPin.DutyCycle(fanOff, cycleLength)
	gpio.StartPwm()

	return nil
}

func SetFanSpeed(percentage uint8) error {
	if percentage < 0 || percentage > 100 {
		return errors.New("Fan speed can't be greater than 100 and less than 0")
	}

	fanControllingPin.DutyCycle(uint32(percentage), cycleLength)
	return nil
}

func CleanUp() {
	gpio.StopPwm()
	gpio.Close()
}
