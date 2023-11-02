package ratelimit

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
)

func Test(t *testing.T) {
	limiter := rate.NewLimiter(1, 1)

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Handle event", i)
		} else {
			fmt.Println("Rate limited", i)
		}
	}
}
