package compare

import "testing"

func TestEqualFloat46(t *testing.T) {
	comparator := EqualFloat64(0.5)

	if comparator(0.5) != true {
		t.Error("Expected EqualFloat64(\"0.5\")(\"0.5\") to equal true")
	}

	if comparator(0.4) != false {
		t.Error("Expected EqualString(\"0.5\")(\"0.4\") to equal false")
	}
}

func TestNotEqualFloat46(t *testing.T) {
	comparator := NotEqualFloat64(0.4)

	if comparator(0.4) != false {
		t.Error("Expected EqualString(\"0.4\")(\"0.4\") to equal false")
	}

	if comparator(0.3) != true {
		t.Error("Expected EqualString(\"0.4\")(\"0.3\") to equal true")
	}
}

func TestLessThanFloat64(t *testing.T) {
	comparator := LessThanFloat64(0.5)

	if comparator(0.4) != true {
		t.Error("Expected LessThanFloat64(0.5)(0.4) to equal true")
	}

	if comparator(0.6) != false {
		t.Error("Expected LessThanUint64(0.5)(0.6) to equal false")
	}

	if comparator(0.6) != false {
		t.Error("Expected LessThanUint64(0.5)(0.6) to equal false")
	}
}

func TestLessThanOrEqualFloat64(t *testing.T) {
	comparator := LessThanOrEqualFloat64(0.5)

	if comparator(0.4) != true {
		t.Error("Expected LessThanOrEqualUint64(0.5)(0.4) to equal true")
	}

	if comparator(0.5) != true {
		t.Error("Expected LessThanOrEqualUint64(0.5)(0.5) to equal true")
	}

	if comparator(0.6) != false {
		t.Error("Expected LessThanOrEqualUint64(0.5)(0.6) to equal false")
	}
}

func TestGreaterThanOrEqualFloat46(t *testing.T) {
	comparator := GreaterThanOrEqualFloat64(0.2)

	if comparator(0.1) != false {
		t.Error("Expected GreaterThanOrEqualUint64(0.2)(0.1) to equal false")
	}

	if comparator(0.2) != true {
		t.Error("Expected GreaterThanOrEqualUint64(0.2)(0.2) to equal true")
	}

	if comparator(0.3) != true {
		t.Error("Expected GreaterThanOrEqualUint64(0.2)(0.3) to equal true")
	}
}

func TestGreaterThanFloat64(t *testing.T) {
	comparator := GreaterThanFloat64(0.2)

	if comparator(0.1) != false {
		t.Error("Expected GreaterThanUint64(0.2)(0.1) to equal false")
	}

	if comparator(0.2) != false {
		t.Error("Expected GreaterThanUint64(0.2)(0.2) to equal false")
	}

	if comparator(0.3) != true {
		t.Error("Expected GreaterThanUint64(0.2)(0.3) to equal true")
	}
}
