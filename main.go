package main

import (
	"math"
	"strconv"

	"github.com/prometheus/common/log"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/gpio"
)

func main() {
	mbot := gobot.NewMaster()
	server := api.NewAPI(mbot)
	// server.AddHandler(api.BasicAuth("gort", "klaatu"))
	server.Port = "3000"
	server.Start()

	mbot.AddCommand("start",
		func(params map[string]interface{}) interface{} {
			fstops, exists := params["fstops"]
			if !exists {
				log.Errorln("No steps especified, ignoring!")
				return nil //TODO: 400?
			}
			var f float64
			var delay float64
			var err error
			switch val := fstops.(type) {
			case string:
				f, err = strconv.ParseFloat(val, 64)
				if err != nil {
					log.Errorf("fstops value can't be converted to a float64")
					return nil
				}
			case float64:
				f = val
			case int:
				f = float64(val)
			}
			d, exists := params["delay"]
			if !exists {
				delay = 0
			} else {
				switch del := d.(type) {
				case string:
					delay, err = strconv.ParseFloat(del, 64)
					if err != nil {
						log.Errorf("delay value can't be converted to a float64")
						return nil
					}
				}
			}
			createTimer(f, delay)
			return "This command is attached to the mcp!"
		})
}

func createTimer(fstop float64, delay float64) {
	log.Infof("Created timer with %f (%f s) starting after %f", fstop, fStop2Second(fstop), delay)
}

func turnOn(led *gpio.LedDriver) {
	err := led.On()
	if err != nil {
		log.Errorf("There was an error turning ON the LED: %e\n", err)
	}
}
func turnOff(led *gpio.LedDriver) {
	err := led.Off()
	if err != nil {
		log.Errorf("There was an error turning OFF the LED: %e\n", err)
	}
}

func fStop2Second(stops float64) float64 {
	return math.Pow(2, stops)
}
