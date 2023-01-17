package middleware

import (

	"time"

)

var Value int = 0

func RateLimitInit() {
	for {
		Value = 0
		time.Sleep(2 * time.Second)
	}
}