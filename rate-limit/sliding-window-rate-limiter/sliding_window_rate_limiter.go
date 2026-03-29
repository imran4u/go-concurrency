package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	swrl := NewSlidingWindowRateLimiter(2*time.Second, 3)
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			if swrl.isAllowed() {
				fmt.Println("Allowed", i)
			} else {
				fmt.Println("Not allowed", i)
			}
			wg.Done()
		}(i)
			time.Sleep(500*time.Millisecond)
	}
	wg.Wait()

}

type SlidingWindowRateLimiter struct {
	duration time.Duration
	limit    int
	request  []time.Time
	mu       sync.Mutex
}

func NewSlidingWindowRateLimiter(duration time.Duration, limit int) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		duration: duration,
		limit:    limit,
		request:  make([]time.Time, 0),
	}
}

func (s *SlidingWindowRateLimiter) isAllowed() bool {
	defer s.mu.Unlock()
	s.mu.Lock()

	now := time.Now()
	cutoff := now.Add(-s.duration)
	//THIS WILL TAKE O(N) TO REMOVE ITEMS
	// iIndex := 0
	// for _, t := range s.request {
	// 	if t.Before(cutoff) {
	// 		iIndex++
	// 		continue
	// 	}
	// 	break
	// }
	// s.request = s.request[iIndex:]

	// remove expired requests (queue style)
	for len(s.request) > 0 && s.request[0].Before(cutoff) {
		s.request = s.request[1:]
	}

	

	if len(s.request) >= s.limit {
		return false
	}
	s.request = append(s.request, now)
	return true
}
