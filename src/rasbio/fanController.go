package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// 控制的引脚（BCM编号）
const fanPin = 17

func main() {
	fmt.Println("fanController is running")
	err := rpio.Open()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Start Fan Controller SuccessFully...")
	}
	defer rpio.Close()

	pin := rpio.Pin(fanPin)
	pin.Input()
	//不断检测cpu温度，当温度超过47度再打开风扇,小于45度则关闭风扇
	//每30s检查一次
	var temp float64
	isHigh := pin.Read() == rpio.High
	pin.Output()
	for {
		temp, err = getTemperature()
		if err != nil {
			fmt.Println(err)
		}
		if temp > 47.0 {
			if !isHigh {
				pin.High()
				fmt.Println("Temperature is too high,Open Fan ")
				isHigh = true
			}
		} else if temp < 45.0 {
			if isHigh {
				pin.Low()
				isHigh = false
				fmt.Println("Fan close")
			}
		}
		time.Sleep(30 * time.Second)
	}
}

// 获取cpu温度
func getTemperature() (float64, error) {
	cmd := exec.Command("vcgencmd", "measure_temp")
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	tempStr := strings.TrimSpace(string(out))
	temps := strings.Split(tempStr, "=")
	temp := temps[1]
	tempFloatStr := temp[:len(temp)-2]
	fmt.Println("温度为:", tempFloatStr)
	return strconv.ParseFloat(tempFloatStr, 64)
}
