package limit

import (
	"context"
	"golang.org/x/sync/semaphore"
	"time"
)

// Limiter provides a way to bound too many concurrent tasks to run.
type Limiter struct {
	sem      *semaphore.Weighted
	duration time.Duration
}

// WaitN waits for n tokens to continue.
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

// New returns a new limiter which ensures that
// at most n tasks can run in the time window of d.
func New(n int64, d time.Duration) *Limiter {
	return &Limiter{
		sem:      semaphore.NewWeighted(n),
		duration: d,
	}
}
