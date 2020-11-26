package ruleSet

import (
	"math/rand"
	"time"
)

const maxJitter = 1000

func doRetry(attempt, maxAttempts int, fn func() (err error, retry bool)) error {
	err, retry := fn()

	attempt++

	if !retry || err == nil || attempt >= maxAttempts {
		return err
	}

	delay := time.Duration(1<<attempt) * time.Millisecond
	delay += time.Duration(rand.Int31n(maxJitter)) * time.Microsecond

	time.Sleep(delay)

	return doRetry(attempt, maxAttempts, fn)
}

func retry(maxAttempts int, fn func() (error, bool)) error {
	return doRetry(0, maxAttempts, fn)
}
