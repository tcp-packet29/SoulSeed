package middleware

import (

	"time"
	"fmt"

)

var Value int = 0

func RateLimitInit() {
	for {
		Value = 0
		time.Sleep(2 * time.Second)
	}
}

func CheckRL() {
	Value++
	if Value > 10 {
		fmt.Println("rate limit exceeded")
		panic("Rate Limit Exceeded")
		
	}
}