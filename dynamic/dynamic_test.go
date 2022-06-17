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

func TestIntegerBreak(t *testing.T) {
	if integerBreak(2) != 1 {
		t.Fatal("failed")
	}
	if integerBreak(10) != 36 {
		t.Fatal("failed")
	}
	if integerBreak2(2) != 1 {
		t.Fatal("failed")
	}
	if integerBreak2(10) != 36 {
		t.Fatal("failed")
	}
}

func TestCanPartition(t *testing.T) {
	if canPartition([]int{1, 5, 11, 5}) != true {
		t.Fatal("failed")
	}
	if canPartition([]int{1, 2, 3, 5}) != false {
		t.Fatal("failed")
	}
}

func TestFindLength(t *testing.T) {
	if findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}) != 3 {
		t.Fatal("failed")
	}
	if findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}) != 5 {
		t.Fatal("failed")
	}
}
