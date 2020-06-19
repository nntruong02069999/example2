package main

import (
	"time"
	"fmt"
	"context"
)
func usingContext () {
	ctx , cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()
	for i := 1 ; i <= 3 ; i++ {
		select {
		case <- ctx.Done() :
			return
		
		default : 
			fmt.Printf("Lần %v\n" , i)
			time.Sleep(3 * time.Second)
		}
	}
}