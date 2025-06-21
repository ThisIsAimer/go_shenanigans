package main

import (
	"fmt"
	"sync"
	"time"
)

type leakyBucket struct{
	mx sync.Mutex
	capacity int
	tokens int
	leakRate time.Duration
	lastLeak time.Time
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *leakyBucket{
	return &leakyBucket{
		capacity: capacity,
		tokens: capacity,
		leakRate: leakRate,
		lastLeak: time.Now(),
	}
}

func (lb *leakyBucket) allow() bool {
	lb.mx.Lock()
	defer lb.mx.Unlock()

	now := time.Now()

	elapsedTime := now.Sub(lb.lastLeak)

	tokensToAdd := int(elapsedTime/lb.leakRate)

	//this is intentional
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity{
		lb.tokens--
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd)* lb.leakRate)

	if lb.tokens > 0{
		lb.tokens--
		return true
	}

	return false

}


func main(){

	myLeakyBucket := NewLeakyBucket(5, time.Second)
	
	for range 20{
		if myLeakyBucket.allow(){
			fmt.Println("access was allowed")
		}else {
			fmt.Println("access denied")
		}

		time.Sleep(time.Second/4)
	}

}