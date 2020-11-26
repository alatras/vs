package ruleSet

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

const maxJitter = 1000

var errMaxAttemptsReached = errors.New("max attempts reached")

func doBackOffRetry(ctx context.Context, attempt, maxAttempts int, initialDelay time.Duration, fn func() (retry bool)) error {
	retry := fn()

	attempt++

	if !retry {
		return nil
	}

	if maxAttempts > 0 && attempt >= maxAttempts {
		return errMaxAttemptsReached
	}

	delay := time.Duration(1<<attempt) * initialDelay
	delay += time.Duration(rand.Int31n(maxJitter)) * time.Microsecond

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		return doBackOffRetry(ctx, attempt, maxAttempts, initialDelay, fn)
	}
}

func backOffRetry(maxAttempts int, initialDelay time.Duration, fn func() bool) error {
	return doBackOffRetry(context.TODO(), 0, maxAttempts, initialDelay, fn)
}

func backOffRetryWithContext(ctx context.Context, initialDelay time.Duration, fn func() bool) error {
	return doBackOffRetry(ctx, 0, 0, initialDelay, fn)
}
