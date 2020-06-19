package main

import (
	"time"
)

func getTimestamp() int {
	dt := time.Now()
	countNamNhuan := soNamNhuan(1970, dt.Year()-1)
	days := dt.YearDay()
	timeStampCurrent := ((dt.Year()-1970)*365+countNamNhuan+days-1)*86400000 + (dt.Hour()-7)*3600000 + dt.Minute()*60000 + dt.Second()*1000 + dt.Nanosecond()/1000000
	return timeStampCurrent
}

func soNamNhuan(start int, end int) int {
	var count int
	for i := end; i >= start; i-- {
		if CheckLeapYear(i) == true {
			count++
		}
	}
	return count
}

func CheckLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
		} else {
			return true
		}
	}
	return false
}
