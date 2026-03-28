package main

import (
	"fmt"
	"sync"
	"time"
)

/**
Downside of fixed-window rate limiter is limit can exceed at the end of first window and starting on new window ( combination of both can form another window but with more call then allowed)
**/

func main() {
	rl := NewRateLimit(2, 2*time.Second)

	fmt.Println(time.Second)
	for range 20 {
		if rl.isAllowed() {
			fmt.Println("Allowed")
		} else {
			fmt.Println("Not allowed")
		}

		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("End")
}

type RateLmit struct {
	limit  int
	window time.Duration

	count     int
	windowEnd time.Time
	mu        sync.Mutex
}

func NewRateLimit(limit int, window time.Duration) *RateLmit {
	now := time.Now()
	windowEnd := now.Add(window)
	return &RateLmit{
		limit:     limit,
		window:    window,
		windowEnd: windowEnd,
	}
}

/*
* steps
 1. first check time and reset value if need
 2. check the count value, if false return.
 3. increment the count value

*
*/
func (r *RateLmit) isAllowed() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.After(r.windowEnd) {
		r.count = 0
		r.windowEnd = now.Add(r.window)
	}

	if r.count >= r.limit {
		return false
	}
	r.count++
	return true

}
