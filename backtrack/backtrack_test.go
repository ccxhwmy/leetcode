package backtrack

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	fmt.Println(permuteWithRecursion([]int{1, 2, 3}))
	fmt.Println(permute([]int{1, 2, 3}))
}

func compareIntSlice(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestRestoreIpAddresses(t *testing.T) {
	if !compareIntSlice(restoreIpAddresses("25525511135"), []string{
		"255.255.11.135","255.255.111.35",
	}) {
		t.Fatal("failed")
	}
	if !compareIntSlice(restoreIpAddresses("0000"), []string{
		"0.0.0.0",
	}) {
		t.Fatal("failed")
	}
	if !compareIntSlice(restoreIpAddresses("101023"), []string{
		"1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3",
	}) {
		t.Fatal("failed")
	}
}
