package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
)

func main() {
	//start(gpio_pin0)
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
	d := DHT11{
		Pin:  rpio.Pin(dataPin26),
		Opin: rpio.Pin(gpioPin19),
	}
	d.DHT11_Detection()
	d.StartDevice()
	var humidity, tempeHigh, tempeLow uint8
	var result0 = d.DHT11_Read_Data(&tempeHigh, &tempeLow, &humidity)
	if result0 == 0 {
		fmt.Printf("DHT11_temp_high = %d\r\n", tempeHigh)
		fmt.Printf("DHT11_temp_low = %d\r\n", tempeLow)
		fmt.Printf("DHT11_humi = %d\r\n", humidity)
	} else {
		fmt.Println("DHT11 DATA Fail \r")
	}
	defer d.CloseDevice()
	d.CloseDevice()

}
