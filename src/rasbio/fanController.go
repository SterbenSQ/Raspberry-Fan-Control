package main

import (
	"flag"
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// 控制的引脚（BCM编号）
//const fanPin = 17

func main() {

	var config struct {
		Pin  int
		Low  float64
		High float64
	}

	flag.IntVar(&config.Pin, "pin", 0, "可编程高压引脚BCM编码")
	flag.Float64Var(&config.Low, "low", 0, "风扇停转的温度")
	flag.Float64Var(&config.High, "high", 0, "风扇开始转动的温度")
	flag.Parse()
	fmt.Printf("获取到配置：config:{pin:%d,low:%f,high:%f}\n", config.Pin, config.Low, config.High)
	if config.Low == 0 || config.High == 0 || config.Pin == 0 {
		flag.Usage()
		return
	}

	//file, err := os.Open("./tsconfig.json")
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("fanController is running")
	err := rpio.Open()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Start Fan Controller SuccessFully...")
	}
	defer rpio.Close()
	fanPin := config.Pin
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
		if temp > config.High {
			if !isHigh {
				pin.High()
				fmt.Println("Temperature is too high,Open Fan ")
				isHigh = true
			}
		} else if temp < config.Low {
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
