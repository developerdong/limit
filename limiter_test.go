package limit

import (
	"context"
	"testing"
	"time"
)

const (
	n        = 60
	duration = 10 * time.Second
)

func TestLimiter_Wait(t *testing.T) {
	limiter := New(n, duration)
	startTime := time.Now()
	for i := 0; i < n+1; i++ {
		_ = limiter.Wait(context.Background())
	}
	if time.Now().Sub(startTime) <= duration {
		t.Error("the number of requests exceeds the limit")
	}
}

func TestLimiter_Pass(t *testing.T) {
	limiter := New(n, duration)
	for i := 0; i < n; i++ {
		if !limiter.Pass() {
			t.Error("the earlier n requests should not fail")
		}
	}
	if limiter.Pass() {
		t.Error("the last request should not succeed")
	}
}
