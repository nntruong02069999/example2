package main

import (
	_ "log"
	_ "time"
	_ "context"
)
type favContextKey string

func main() {
	//bai1()
	//bai2
	// curentTime := getTimestamp()
	// log.Printf("Timestamp : %v " ,curentTime)

	// Bai 3
	//usingContext()

	// Bai 4
	// calTime := CaculateMinutesOfTime(1592190294764144364)
	// log.Printf("Số phút sau khi quy đổi : %v phút", calTime)
	
	// Bai 5
	caculateDayOfWeek(1592190385)

	// Bai 7 
	// key := favContextKey("timeCurrent")
	// ctx := context.WithValue(context.Background(), key, time.Now().UnixNano())
	// timeBetween := bai7(ctx, key)
	// log.Printf("Độ chênh lệch thời gian : %v ns", timeBetween)
}
