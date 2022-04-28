package dynamic

import "testing"

func TestLongestParlindrome(t *testing.T) {
	res := longestPalindrome("babad")
	if res != "bab" && res != "aba" {
		t.Fatal("failed")
	}
	if longestPalindrome("cbbd") != "bb" {
		t.Fatal("failed")
	}
	res = longestPalindrome2("babad")
	if res != "bab" && res != "aba" {
		t.Fatal("failed")
	}
	if longestPalindrome2("cbbd") != "bb" {
		t.Fatal("failed")
	}
}
