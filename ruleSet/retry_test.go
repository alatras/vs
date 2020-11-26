package ruleSet

import (
	"context"
	"errors"
	"testing"
	"time"
)

const testDelay = 1 * time.Microsecond

func TestRetry_MaxAttempts(t *testing.T) {
	attempts := 5
	count := 0

	var err error

	retryErr := backOffRetry(attempts, testDelay, func() (retry bool) {
		count++
		err = errors.New("error")
		return true
	})

	if retryErr != errMaxAttemptsReached {
		t.Fatalf("Retry should return '%s' error but got '%s'", errMaxAttemptsReached, retryErr)
	}

	if err == nil {
		t.Fatal("Error was not returned")
	}

	if err.Error() != "error" {
		t.Fatal("Different error returned")
	}

	if count != attempts {
		t.Fatalf("Should have %d attempts but got %d", attempts, count)
	}
}

func TestRetry_SuccessAfterFailure(t *testing.T) {
	attempts := 5
	recoverOnAttempt := 3
	count := 0

	var err error

	retryErr := backOffRetry(attempts, testDelay, func() (retry bool) {
		count++

		if count == recoverOnAttempt {
			return false
		}

		return true
	})

	if retryErr != nil {
		t.Fatalf("Retry should return no error but got '%s'", retryErr)
	}

	if err != nil {
		t.Fatal("No error should be returned")
	}

	if count != recoverOnAttempt {
		t.Fatalf("Should recover after %d attempts but got %d", recoverOnAttempt, count)
	}
}

func TestRetry_NoError(t *testing.T) {
	attempts := 5
	count := 0

	retryErr := backOffRetry(attempts, testDelay, func() (retry bool) {
		count++
		return false
	})

	if retryErr != nil {
		t.Fatalf("Retry should return no error but got '%s'", retryErr)
	}

	if count != 1 {
		t.Fatalf("Should return immediately after first attempt but had %d attempts", count)
	}
}

func TestRetry_PermanentError(t *testing.T) {
	attempts := 5
	count := 0

	msg := "Permanent error"

	var err error

	retryErr := backOffRetry(attempts, testDelay, func() (retry bool) {
		count++
		err = errors.New(msg)
		return false
	})

	if retryErr != nil {
		t.Fatalf("Retry should return no error but got '%s'", retryErr)
	}

	if err == nil {
		t.Fatal("Error should be returned")
	}

	if err.Error() != msg {
		t.Fatalf("Error message should be '%s' but got '%s'", msg, err)
	}

	if count != 1 {
		t.Fatalf("Should return immediately after first attempt but had %d attempts", count)
	}
}

func TestRetry_WithTimeout(t *testing.T) {
	count := 0

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Millisecond)
	defer cancel()

	retryErr := backOffRetryWithContext(ctx, testDelay, func() (retry bool) {
		count++
		return true
	})

	expectErr := context.DeadlineExceeded

	if retryErr != expectErr {
		t.Fatalf("Error message should be '%s' but got '%s'", expectErr, retryErr)
	}
}
