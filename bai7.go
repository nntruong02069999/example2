package main

import (
	"context"
	_ "strconv"
	"time"
)

func bai7(ctx context.Context, key favContextKey) int64 {
	val := (ctx.Value(key).(int64))
	timeBefore := val
	var timeNow int64
	select {
	case <- time.After(3 * time.Second) :
		timeNow = time.Now().UnixNano()
	}
	return (timeNow - timeBefore)
}
