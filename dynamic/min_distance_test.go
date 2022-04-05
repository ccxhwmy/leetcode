package dynamic

import "testing"

func TestMinDistance(t *testing.T) {
	if minDistance("horse", "ros") != 3 {
		t.Fatal("failed")
	}
	if minDistance("intention", "execution") != 5 {
		t.Fatal("failed")
	}
}
