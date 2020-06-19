package main

import (
	"fmt"
	"time"
)

func vongLap() {
	fmt.Println("time now: {milliseconds}")
	time.Sleep(3 * time.Second)
}


func bai1 () {
	for i := 0; i< 3 ; i++ {
		vongLap()
	}
	fmt.Println("Kết thúc")
}