package limit

import (
	"context"
	"golang.org/x/sync/semaphore"
	"time"
)

// Limiter provides a way to bound too many concurrent tasks to run. It is like a
// combination of sliding window and token bucket.
type Limiter struct {
	sem      *semaphore.Weighted
	duration time.Duration
}

// WaitN waits for n tokens to continue, blocking until resources are available
// or ctx is done. On success, returns nil. On failure, returns ctx.Err() and
// leaves the limiter unchanged. If ctx is already done, WaitN may still succeed
// without blocking.
func (l *Limiter) WaitN(ctx context.Context, n int64) error {
	var err error
	if err = l.sem.Acquire(ctx, n); err == nil {
		time.AfterFunc(l.duration, func() {
			l.sem.Release(n)
		})
	}
	return err
}

// Wait waits for 1 token to continue.
func (l *Limiter) Wait(ctx context.Context) error {
	return l.WaitN(ctx, 1)
}

// PassN acquires n tokens to continue without blocking. On success, returns
// true. On failure, returns false and leaves the limiter unchanged.
func (l *Limiter) PassN(n int64) bool {
	var success bool
	if success = l.sem.TryAcquire(n); success {
		time.AfterFunc(l.duration, func() {
			l.sem.Release(n)
		})
	}
	return success
}

// Pass acquires 1 token to continue without blocking.
func (l *Limiter) Pass() bool {
	return l.PassN(1)
}

// New returns a new limiter which ensures that at most n tasks can run in the
// time window of d.
func New(n int64, d time.Duration) *Limiter {
	return &Limiter{
		sem:      semaphore.NewWeighted(n),
		duration: d,
	}
}
