# fan-controller

fan-controller is a simple and minimalistic app which allows to control Raspberry Pi fan speed
using built-in [hardware PWM](https://www.raspberrypi.org/documentation/usage/gpio/) capabilities.

## How to use?

First - you need to create a simple curquit which uses signal from *GPIO 18* pin to control 5v output power. 
I use 2n2222 transistor for this purpose. To protect *GPIO 18* from high current I reccomend to connect transistor base to the pin using 1kOm resistor. 

### Circuit 
![circuit](https://raw.githubusercontent.com/AndrewKovalenko/fan-controller/master/readme-images/circuit_fan.png?token=AAPEEKLWRXFZXBEVLUCFZKS637OWK)


## Build
In order to build *fan-controller* you need to have latest version of [Go](https://golang.org/doc/install) installled on your computer.

Once you have *Go* installed and `go version` command doesn't return you an error - clone this repository and run `build.sh` script project directory.
You'll find a distribution package in a `build` folder.

### Distrubution package structure

There are 3 files in a *fan-controller* distribution package:
* **fan-controller** - a binary file of the application
* **fan-controller-config.yaml** - a config file
* **start.sh** - optional start script 

## Installation 
1. Copy all 3 files to a folder on your Raspberry Pi
2. Run `sudo nano /etc/rc.local`
3. Add `<path-to-the-folder-with-fan-controller>/start.sh` before `exit 0`
4. Save
5. Reboot

## Config
```
temperatureCheckFrequency: 20
stepdownTemperature: 3
fanSpeedSettings:
  50: 50
  54: 70
  58: 100
```

Config has 3 parameters:
`temperatureCheckFrequency` - a number of seconds the app will "wait" before checking the temperature again
`stepdownTemperature` - a number of degrees **below** current temperature setting the fan will keep running with current speed setting. 
For example: once CPU temperature reaches 50 deg - the fan will start running at 50% of it's speed. If if stepdown temperature is set to 3 deg - the fun will keep running at 50% until CPU cools down to 47 deg.
`fanSpeedSettings` - a key-value collection of <CPU temperature>:<fan speed>. CPU temperature is measured in Celsius. Fan Speed is measured in percents. 
