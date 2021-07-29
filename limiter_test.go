package simple_rate_limiter

import (
	"context"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	n := 60
	duration := 10 * time.Second
	limiter := NewLimiter(int64(n), duration)
	for i := 0; i < n+1; i++ {
		_ = limiter.Wait(context.Background())
	}
}
