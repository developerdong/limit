package limit

import (
	"context"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	n := 60
	duration := 10 * time.Second
	limiter := New(int64(n), duration)
	startTime := time.Now()
	for i := 0; i < n+1; i++ {
		_ = limiter.Wait(context.Background())
	}
	if time.Now().Sub(startTime) <= duration {
		t.Error("the number of requests exceeds the limit")
	}
}
