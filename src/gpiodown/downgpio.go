package main

import (
	"flag"
	"fmt"
	"github.com/SterbenSQ/go-rpio/v4"
)

//const activePin19 = 19

//const activePin26 = 26

func main() {
	rpio.Open()
	defer rpio.Close()

	var pin uint

	flag.UintVar(&pin, "pin", 0, "要停用的引脚编号（BCM）")
	flag.Parse()

	if pin == 0 {
		fmt.Println("请输入引脚编号（BCM）")
		return
	}

	pind := uint8(pin)
	fmt.Println("pin bcm code:", pin)
	pin19 := rpio.Pin(pind)
	pin19.Output()
	pin19.Low()
}
