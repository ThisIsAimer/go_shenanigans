package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func newRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{tokens: make(chan struct{}, rateLimit), refillTime: refillTime}

	//this fills up tokens
	for range rateLimit {
		// struct{}{} takes up 0 bytes of memory
		rl.tokens <- struct{}{}
	}

	go rl.startRefill()

	return rl
}

func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	//this adds tokens every tick if token is used
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		default:
			fmt.Print("")
		}
	}
}

func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {

	rateLimiter := newRateLimiter(5, time.Second)

	//this consumes tokens
	for range 20 {
		if rateLimiter.allow() {
			fmt.Println("request allowed")

		} else {
			fmt.Println("request denied")
		}

		time.Sleep(time.Millisecond * 300)
	}

}
