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
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	if numIslands(grid1) != 1 {
		t.Fatal("failed")
	}
}

func TestSpiralMatrix(t *testing.T) {
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	t.Log(spiralMatrix(matrix1))
	t.Log(spiralMatrix(matrix2))
}

func TestLengthOfLIS(t *testing.T) {
	if lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}) != 4 {
		t.Fatal("failed")
	}
	if lengthOfLIS([]int{0, 1, 0, 3, 2, 3}) != 4 {
		t.Fatal("failed")
	}
	if lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}) != 1 {
		t.Fatal("failed")
	}
}

func TestCoinChange(t *testing.T) {
	if coinChange([]int{1, 2, 5}, 11) != 3 {
		t.Fatal("failed")
	}
}
