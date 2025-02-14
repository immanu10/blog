Title: Rate Limiting
Date: 14-Feb-2025

Rate limiting is a crucial mechanism in system design used to control the rate at which clients can access an API or service.  It prevents abuse, protects against denial-of-service (DoS) attacks, and ensures fair usage for all users.  By throttling excessive requests, rate limiting helps maintain service availability and prevents overload.

There are several common algorithms used to implement rate limiting:

* **Token Bucket:**  Imagine a bucket that's filled with tokens at a fixed rate.  Each request requires a token to be processed. If a request arrives and there's a token available, the request is processed, and the token is removed. If no tokens are available, the request is rejected.  This algorithm allows for bursts of traffic as long as tokens are available in the bucket.

* **Leaky Bucket:** Similar to the token bucket, but instead of discrete tokens, the leaky bucket holds requests.  Requests are added to the bucket at the arrival rate. The bucket "leaks" requests at a fixed processing rate.  If the bucket overflows, incoming requests are discarded. This algorithm smooths out bursts by enforcing a constant processing rate.

* **Fixed Window Counter:** A counter is maintained for a fixed time window (e.g., 1 minute).  Each request increments the counter. If the counter exceeds a predefined threshold within the time window, subsequent requests are rejected until the next window begins.  This algorithm is simple but can be susceptible to bursts at the boundaries of time windows.

* **Sliding Window Counter:** A more sophisticated version of the fixed window counter.  It divides the time window into smaller sub-windows and keeps track of the count for each sub-window. The total count for the current window is calculated by summing the counts of the relevant sub-windows.  This approach provides smoother rate limiting compared to the fixed window counter.


Here's a simplified example of rate limiting using the token bucket algorithm in Go:

```go
package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	bucket      chan struct{}
	rate        int
	capacity    int
	lastRefill  time.Time
	refillRate time.Duration
}

func NewRateLimiter(rate int, capacity int) *RateLimiter {
	limiter := &RateLimiter{
		bucket:      make(chan struct{}, capacity),
		rate:        rate,
		capacity:    capacity,
		lastRefill:  time.Time{},
		refillRate: time.Second / time.Duration(rate),
	}
	limiter.fillBucket() // Initial fill
	go limiter.refillLoop()
	return limiter
}


func (rl *RateLimiter) Allow() bool {

	if rl.bucket <- struct{}{}:
		return true
	}
	return false
}


func (rl *RateLimiter) fillBucket() {
	now := time.Now()
	timeElapsed := now.Sub(rl.lastRefill)
	tokensToAdd := int(timeElapsed / rl.refillRate)
	if tokensToAdd > 0 {
		for i := 0; i < tokensToAdd && len(rl.bucket) < rl.capacity; i++ {
			rl.bucket <- struct{}{}
		}
		rl.lastRefill = now

	}
}

func (rl *RateLimiter) refillLoop() {

	for {
		rl.fillBucket()
		time.Sleep(rl.refillRate)
	}
}


func main() {
	limiter := NewRateLimiter(2, 5) // 2 requests per second, bucket size 5

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Request processed", i)
		} else {
			fmt.Println("Request throttled", i)

		}
		time.Sleep(200 * time.Millisecond)
	}

}

```

This example demonstrates a basic implementation. In a real-world scenario, you would likely integrate rate limiting with your API gateway or use a dedicated service like Redis for distributed rate limiting.  The choice of algorithm and implementation depends on the specific requirements of your application and the desired level of control over request rates.
