package compare

import "testing"

func TestEqualString(t *testing.T) {
	comparator := EqualString("foo")

	if comparator("foo") != true {
		t.Error("Expected EqualString(\"foo\")(\"foo\") to equal true")
	}

	if comparator("bar") != false {
		t.Error("Expected EqualString(\"foo\")(\"bar\") to equal false")
	}
}

func TestNotEqualString(t *testing.T) {
	comparator := NotEqualString("foo")

	if comparator("foo") != false {
		t.Error("Expected EqualString(\"foo\")(\"foo\") to equal false")
	}

	if comparator("bar") != true {
		t.Error("Expected EqualString(\"foo\")(\"bar\") to equal true")
	}
}
