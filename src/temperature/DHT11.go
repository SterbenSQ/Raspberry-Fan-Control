package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

// const gpioPin19 uint8 = 19
const dataPin26 uint8 = 26

type DHT11 struct {
	Pin rpio.Pin
	//Opin rpio.Pin
}

// 为传感器供电
func (d *DHT11) StartDevice() {
	//d.Opin.Output()
	//d.Opin.Write(rpio.High)
}

// 关闭传感器
func (d *DHT11) CloseDevice() {
	//d.Opin.Low()
	d.Pin.Low()
}

func (d *DHT11) Dht11Detection() {
	if d.Dht11Init() == 0 {
		Dht11Flag := 1
		fmt.Println("DHT11 OK \r\n", Dht11Flag)
	} else {
		fmt.Println("DHT11 Fail \r\n")
	}
}

func (d *DHT11) Dht11Init() uint8 {
	time.Sleep(1 * time.Second)
	d.Dht11Rst()
	return d.Dht11Check()
}

/*
*

	读取传感器1字节
*/
func (d *DHT11) Dht11ReadByte() uint8 {
	var i, dat uint8
	dat = 0x00
	for i = 0; i < 8; i++ {
		dat <<= 1
		dat |= d.Dht11ReadBit()
	}
	return dat
}

// 读取温度器返回的结果-读取一位
func (d *DHT11) Dht11ReadBit() uint8 {
	var retry uint8
	retry = 0
	for (d.Pin.Read() == rpio.High) && retry < 100 { //等待变为低电平
		retry++
		time.Sleep(2 * time.Microsecond)
	}

	for (d.Pin.Read() == rpio.Low) && retry < 100 { //等待变高电平
		retry++
		time.Sleep(2 * time.Microsecond)
	}
	time.Sleep(20 * time.Microsecond)
	time.Sleep(20 * time.Microsecond)

	if d.Pin.Read() == rpio.High { //用于判断高低电平，即数据1或0
		return 1
	} else {
		return 0
	}
}

func (d *DHT11) Dht11Check() uint8 {
	var retry uint8 = 0
	d.Pin.Input()
	for d.Pin.Read() == rpio.High && retry < 100 { //DHT11会拉低40~80us
		retry++
		time.Sleep(2 * time.Microsecond)
	}
	if retry >= 100 {
		return 1
	} else {
		retry = 0
	}
	for d.Pin.Read() == rpio.Low && retry < 100 { //DHT11拉低后会再次拉高40~80us
		retry++
		time.Sleep(2 * time.Microsecond)
	}
	if retry >= 100 {
		return 1
	}
	return 0
}

// 复位
func (d *DHT11) Dht11Rst() {
	d.Pin.Output()
	d.Pin.Low()
	time.Sleep(25 * time.Millisecond)
	d.Pin.High()
	time.Sleep(20 * time.Microsecond)
}

// 获取传感器温湿度
func (d *DHT11) Dht11ReadData(temphigh *uint8, templow *uint8, humi *uint8) uint8 {

	var buf = make([]uint8, 5)
	d.Dht11Rst() //DHT11端口复位，发出起始信号
	defer func() {
		d.CloseDevice()
	}()
	if d.Dht11Check() == 0 { //等待DHT11回应，0为成功回应
		for i := 0; i < 5; i++ { //读取40位数据
			buf[i] = d.Dht11ReadByte() //读出数据
			fmt.Printf("DHT11 Read Byte : %d\r\n", buf[i])
		}
		if (buf[0] + buf[1] + buf[2] + buf[3]) == buf[4] { //数据校验

			*humi = buf[0]     //将湿度整数值放入指针humi
			*temphigh = buf[2] //将温度整数值放入指针temphigh
			*templow = buf[3]  //将温度小数值放入指针templow
		}
	} else {
		return 1
	}
	return 0
}
