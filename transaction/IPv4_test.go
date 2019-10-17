package transaction

import "testing"

func TestIsIPv4(t *testing.T) {
	if ok := IsIPv4("0.0.0.0"); !ok {
		t.Error("Valid IP address should have passed assertion")
	}

	if ok := IsIPv4("255.255.255.255"); !ok {
		t.Error("Valid IP address should have passed assertion")
	}

	if ok := IsIPv4("256.256.256.256"); ok {
		t.Error("Invalid IP address should have passed assertion")
	}

	if ok := IsIPv4("a.b.c.d"); ok {
		t.Error("Invalid IP address should have passed assertion")
	}

	if ok := IsIPv4("0.0.0.0.0"); ok {
		t.Error("Invalid IP address should have passed assertion")
	}

	if ok := IsIPv4("0.0.0"); ok {
		t.Error("Invalid IP address should have passed assertion")
	}

	if ok := IsIPv4("0.0"); ok {
		t.Error("Invalid IP address should have passed assertion")
	}

	if ok := IsIPv4("0"); ok {
		t.Error("Invalid IP address should have passed assertion")
	}
}
