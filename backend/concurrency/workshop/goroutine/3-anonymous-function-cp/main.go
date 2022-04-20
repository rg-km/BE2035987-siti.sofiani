package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		result := 0
		for i := 0; i < 5; i++ {
			result += i
		}
		time.Sleep(10 * time.Millisecond)
		fmt.Println("main stop")
	}()
}
