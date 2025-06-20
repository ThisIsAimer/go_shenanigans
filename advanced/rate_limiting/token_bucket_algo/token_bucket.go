package main

import (
	"fmt"
	"time"
)

type RateLimiter struct{
	tokens chan struct{}
	refillTime time.Duration
}


func newRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{ tokens: make(chan struct{}, rateLimit), refillTime: refillTime }

	for range rateLimit{
		// struct{}{} takes up 0 bytes of memory
		rl.tokens <- struct{}{}
	}

	go rl.startRefill()

	return rl
}

func (rl *RateLimiter) startRefill(){
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for {
		select{
		case <- ticker.C:
			select{
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (rl *RateLimiter) allow() bool{
	select{
	case <- rl.tokens:
		return true
	default:
		return false
	}
}


func main(){


	rateLimiter := newRateLimiter(5, time.Second)


	for range 10{
		if rateLimiter.allow(){
			fmt.Println("request allowed")
			
		}else{
			fmt.Println("request denied")
		}

		time.Sleep(time.Microsecond * 200)
	}

}