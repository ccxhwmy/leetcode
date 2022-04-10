package stack

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	if longestValidParentheses(`(()`) != 2 {
		t.Fatal("failed")
	}
	if longestValidParentheses(`)()())`) != 4 {
		t.Fatal("failed")
	}
	if longestValidParentheses("") != 0 {
		t.Fatal("failed")
	}
}
