package main

import (
	"log"
	"math"
	"net/http"

	"github.com/warthog618/gpiod"

	"github.com/Willyfrog/red-time/server"
)

func main() {
	c, _ := gpiod.NewChip("gpiochip0", gpiod.WithConsumer("softwire"))
	Port := 3000 // TODO: accept a flag to setup the port
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
	log.Printf("Server will be running in post %d", Port)
	defer c.Close()
}

func createTimer(fstop float64, delay float64) {
	log.Printf("Created timer with %f (%f s) starting after %f", fstop, fStop2Second(fstop), delay)
}

func turnOn(led *gpiod.Line) {
	err := led.SetValue(1)
	if err != nil {
		log.Printf("There was an error turning ON the LED: %e\n", err)
	}
}
func turnOff(led *gpiod.Line) {
	err := led.SetValue(0)
	if err != nil {
		log.Printf("There was an error turning OFF the LED: %e\n", err)
	}
}

func fStop2Second(stops float64) float64 {
	return math.Pow(2, stops)
}
