package application

import (
	"fmt"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

func InitializeFanController() {
	fmt.Println("Controller is initialized")
	rpio.StartPwm()
}
