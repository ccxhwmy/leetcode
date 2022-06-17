package other

import "testing"

func TestAddStrings(t *testing.T) {
	if addStrings("12345", "11111") != "23456" {
		t.Log("failed")
	}
}

func TestMultiply(t *testing.T) {
	if multiply("123", "456") != "56088" {
		t.Log("failed")
	}
}

func TestRand10(t *testing.T) {
	t.Log(rand10())
}

func TestRotate(t *testing.T) {
	testArr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(testArr)
	t.Log(testArr)
}
