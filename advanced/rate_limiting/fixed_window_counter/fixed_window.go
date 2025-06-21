package main

import (
	"fmt"
	"sync"
	"time"
)

// In fixed window algorythm we divide time into fixed windows where
// limited number of requests can be accessed.

type rateLimiter struct{
	mu sync.Mutex
	count int
	limit int
	window time.Duration
	resetTime time.Time
}

func newRateLimiter (limit int, window time.Duration) *rateLimiter{
return &rateLimiter{
	limit: limit,
	window: window,
}
}

func (rl *rateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	//its going to run first time and make increase reset time
	if now.After(rl.resetTime){
		rl.resetTime = now.Add(rl.window)
		rl.count = 0 
	}

	// going to increase till limit then deny access
	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}


func main(){

	myLimiter := newRateLimiter(5, time.Second)
	
	for range 20{
		go func(){
			if myLimiter.Allow(){
				fmt.Println("access allowed")
			} else{
				fmt.Println("access denied")
			}
		}()
		time.Sleep(time.Second/6)
	}

}