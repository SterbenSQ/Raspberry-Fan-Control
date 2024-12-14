package main

import (
	"flag"
	"fmt"
	"github.com/SterbenSQ/go-rpio/v4"
	"log"
)

//const activePin19 = 19
//const activePin26 = 26

func main() {

	var pin int64

	flag.Int64Var(&pin, "pin", 0, "要启用的引脚编号（BCM）")
	flag.Parse()

	if pin == 0 {
		fmt.Println("请输入引脚编号（BCM）")
		return
	}

	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := rpio.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	pin19 := rpio.Pin(pin)
	//pin26 := rpio.Pin(activePin26)
	//pin26.Output()
	//pin26.Low()
	pin19.Output()
	pin19.High()
	err = rpio.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}
