package other

import "testing"

func TestAddStrings(t *testing.T) {
	if addStrings("12345", "11111") != "23456" {
		t.Log("failed")
	}
}

func TestMultiply(t *testing.T) {
	if multiply("123", "456") != "56088" {
		t.Log("failed")
	}
}
