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

func TestNumIslands(t *testing.T) {
	grid1 := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}
	if numIslands(grid1) != 1 {
		t.Fatal("failed")
	}
}
