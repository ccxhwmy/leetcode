package slide_win

import "testing"

func TestMinWindow(t *testing.T) {
	if minWindow("ADOBECODEBANC", "ABC") != "BANC" {
		t.Fatal("failed")
	}
	if minWindow("a", "a") != "a" {
		t.Fatal("failed")
	}
	if minWindow("a", "aa") != "" {
		t.Fatal("failed")
	}
}
