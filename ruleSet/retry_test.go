package ruleSet

import (
	"errors"
	"testing"
)

func TestRetry_Attempts(t *testing.T) {
	attempts := 5
	count := 0

	err := retry(attempts, func() (err error, retry bool) {
		count++
		return errors.New("error"), true
	})

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

	err := retry(attempts, func() (err error, retry bool) {
		count++

		if count == recoverOnAttempt {
			return nil, false
		}

		return errors.New("error"), true
	})

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

	err := retry(attempts, func() (err error, retry bool) {
		count++
		return nil, false
	})

	if err != nil {
		t.Fatal("No error should be returned")
	}

	if count != 1 {
		t.Fatalf("Should return immediately after first attempt but had %d attempts", count)
	}
}

func TestRetry_PermanentError(t *testing.T) {
	attempts := 5
	count := 0

	msg := "Permanent error"

	err := retry(attempts, func() (err error, retry bool) {
		count++
		return errors.New(msg), false
	})

	if err == nil {
		t.Fatal("Error should be returned")
	}

	if err.Error() != msg {
		t.Fatalf("Error message should be '%s' but got '%s'", msg, err.Error())
	}

	if count != 1 {
		t.Fatalf("Should return immediately after first attempt but had %d attempts", count)
	}
}
