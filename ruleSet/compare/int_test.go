package compare

import "testing"

func TestLessThanInt(t *testing.T) {
	comparator := LessThanInt(2)

	if comparator(1) != true {
		t.Error("Expected LessThanInt(2)(1) to equal true")
	}

	if comparator(2) != false {
		t.Error("Expected LessThanInt(2)(2) to equal false")
	}

	if comparator(3) != false {
		t.Error("Expected LessThanInt(2)(3) to equal false")
	}
}

func TestLessThanOrEqualInt(t *testing.T) {
	comparator := LessThanOrEqualInt(2)

	if comparator(1) != true {
		t.Error("Expected LessThanOrEqualInt(2)(1) to equal true")
	}

	if comparator(2) != true {
		t.Error("Expected LessThanOrEqualInt(2)(2) to equal true")
	}

	if comparator(3) != false {
		t.Error("Expected LessThanOrEqualInt(2)(3) to equal false")
	}
}

func TestEqualInt(t *testing.T) {
	comparator := EqualInt(2)

	if comparator(1) != false {
		t.Error("Expected EqualInt(2)(1) to equal false")
	}

	if comparator(2) != true {
		t.Error("Expected EqualInt(2)(2) to equal true")
	}

	if comparator(3) != false {
		t.Error("Expected EqualInt(2)(3) to equal false")
	}
}

func TestNotEqualInt(t *testing.T) {
	comparator := NotEqualInt(2)

	if comparator(1) != true {
		t.Error("Expected NotEqualInt(2)(1) to equal true")
	}

	if comparator(2) != false {
		t.Error("Expected NotEqualInt(2)(2) to equal false")
	}

	if comparator(3) != true {
		t.Error("Expected NotEqualInt(2)(3) to equal true")
	}
}

func TestGreaterThanOrEqualInt(t *testing.T) {
	comparator := GreaterThanOrEqualInt(2)

	if comparator(1) != false {
		t.Error("Expected GreaterThanOrEqualInt(2)(1) to equal false")
	}

	if comparator(2) != true {
		t.Error("Expected GreaterThanOrEqualInt(2)(2) to equal true")
	}

	if comparator(3) != true {
		t.Error("Expected GreaterThanOrEqualInt(2)(3) to equal true")
	}
}

func TestGreaterThanInt(t *testing.T) {
	comparator := GreaterThanInt(2)

	if comparator(1) != false {
		t.Error("Expected GreaterThanInt(2)(1) to equal false")
	}

	if comparator(2) != false {
		t.Error("Expected GreaterThanInt(2)(2) to equal false")
	}

	if comparator(3) != true {
		t.Error("Expected GreaterThanInt(2)(3) to equal true")
	}
}
