package main

import (
	"fmt"
	"github.com/SterbenSQ/go-rpio/v4"
	"log"
)

func main() {
	//start(gpio_pin0)
	err := rpio.Open()
	defer rpio.Close()
	if err != nil {
		log.Fatal(err)
	}
	d := DHT11{
		Pin:  rpio.Pin(dataPin26),
		Opin: rpio.Pin(gpioPin19),
	}
	d.Dht11Detection()
	d.StartDevice()
	var humidity, tempeHigh, tempeLow uint8
	var result0 = d.Dht11ReadData(&tempeHigh, &tempeLow, &humidity)
	if result0 == 0 {
		fmt.Printf("DHT11_temp_high = %d\r\n", tempeHigh)
		fmt.Printf("DHT11_temp_low = %d\r\n", tempeLow)
		fmt.Printf("DHT11_humi = %d\r\n", humidity)
	} else {
		fmt.Println("DHT11 DATA Fail \r")
	}
	d.CloseDevice()
}
