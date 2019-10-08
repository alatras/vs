package compare

import "testing"

func TestLessThanInt(t *testing.T) {
	comparator := LessThanUint64(2)

	if comparator(1) != true {
		t.Error("Expected LessThanUint64(2)(1) to equal true")
	}

	if comparator(2) != false {
		t.Error("Expected LessThanUint64(2)(2) to equal false")
	}

	if comparator(3) != false {
		t.Error("Expected LessThanUint64(2)(3) to equal false")
	}
}

func TestLessThanOrEqualInt(t *testing.T) {
	comparator := LessThanOrEqualUint64(2)

	if comparator(1) != true {
		t.Error("Expected LessThanOrEqualUint64(2)(1) to equal true")
	}

	if comparator(2) != true {
		t.Error("Expected LessThanOrEqualUint64(2)(2) to equal true")
	}

	if comparator(3) != false {
		t.Error("Expected LessThanOrEqualUint64(2)(3) to equal false")
	}
}

func TestEqualInt(t *testing.T) {
	comparator := EqualUint64(2)

	if comparator(1) != false {
		t.Error("Expected EqualUint64(2)(1) to equal false")
	}

	if comparator(2) != true {
		t.Error("Expected EqualUint64(2)(2) to equal true")
	}

	if comparator(3) != false {
		t.Error("Expected EqualUint64(2)(3) to equal false")
	}
}

func TestNotEqualInt(t *testing.T) {
	comparator := NotEqualUint64(2)

	if comparator(1) != true {
		t.Error("Expected NotEqualUint64(2)(1) to equal true")
	}

	if comparator(2) != false {
		t.Error("Expected NotEqualUint64(2)(2) to equal false")
	}

	if comparator(3) != true {
		t.Error("Expected NotEqualUint64(2)(3) to equal true")
	}
}

func TestGreaterThanOrEqualInt(t *testing.T) {
	comparator := GreaterThanOrEqualUint64(2)

	if comparator(1) != false {
		t.Error("Expected GreaterThanOrEqualUint64(2)(1) to equal false")
	}

	if comparator(2) != true {
		t.Error("Expected GreaterThanOrEqualUint64(2)(2) to equal true")
	}

	if comparator(3) != true {
		t.Error("Expected GreaterThanOrEqualUint64(2)(3) to equal true")
	}
}

func TestGreaterThanInt(t *testing.T) {
	comparator := GreaterThanUint64(2)

	if comparator(1) != false {
		t.Error("Expected GreaterThanUint64(2)(1) to equal false")
	}

	if comparator(2) != false {
		t.Error("Expected GreaterThanUint64(2)(2) to equal false")
	}

	if comparator(3) != true {
		t.Error("Expected GreaterThanUint64(2)(3) to equal true")
	}
}
