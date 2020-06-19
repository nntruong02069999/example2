package main


func CaculateMinutesOfTime(timeNanoSecond int) int {
	minutesOfTime := timeNanoSecond / 60000000
	return minutesOfTime
} 