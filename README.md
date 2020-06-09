# fan-controller

fan-controller is a simple and minimalistic app which allows to control Raspberry Pi fan speed
using built-in [hardware PWM](https://www.raspberrypi.org/documentation/usage/gpio/) capabilities.

## How to use?

First - you need to create a simple curquit which uses signal from *GPIO 18* pin to control 5v output power. 
I use 2n2222 transistor for this purpose. To protect *GPIO 18* from high current I reccomend to connect transistor base to the pin using 1kOm resistor. 

### Cirquit 
![Fan Switch circuit]()
