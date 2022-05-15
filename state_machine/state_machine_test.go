package state_machine

import "testing"

func TestMyAtoi(t *testing.T) {
	res := myAtoi("-91283472332")
	if res != -2147483648 {
		t.Fatal("failed")
	}
	if myAtoi("words and 987") != 0 {
		t.Fatal("failed")
	}
	if myAtoi("42") != 42 {
		t.Fatal("failed")
	}
	if myAtoi("-42") != -42 {
		t.Fatal("failed")
	}
	if myAtoi("4193 with words") != 4193 {
		t.Fatal("failed")
	}
}
