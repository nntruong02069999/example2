package main

import (
	_ "fmt"
	"log"
	"time"
)

func caculateDayOfWeek(inputTime int64){
	timeConverted := time.Time(time.Unix(0,inputTime * int64(time.Millisecond)))
	log.Printf("Ngày trong tuần : %v", timeConverted.Weekday() )
	log.Printf("Số thứ tự trong tuần : %v", int(timeConverted.Weekday()))

}