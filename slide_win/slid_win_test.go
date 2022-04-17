package slide_win

import (
	"leetcode/base"
	"testing"
)

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

func TestMinSlideWindow1(t *testing.T) {
	res := maxSlidingWindowWithPriorityQueue([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	if !base.CompareArrayOfInt(res, []int{3, 3, 5, 5, 6, 7}) {
		t.Fatal("failed")
	}
}

func TestLongestSubstrWithoutRepeatChar(t *testing.T) {
	if longestSubStrWithoutRepeatChar("abcabcbb") != 3 {
		t.Fatal("failed")
	}
	if longestSubStrWithoutRepeatChar("bbbbb") != 1 {
		t.Fatal("failed")
	}
	if longestSubStrWithoutRepeatChar("pwwkew") != 3 {
		t.Fatal("failed")
	}
}
