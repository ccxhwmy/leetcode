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
