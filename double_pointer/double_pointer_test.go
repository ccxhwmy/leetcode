package double_pointer

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}

func TestCompareVersion(t *testing.T) {
	if compareVersion("1.01", "1.001") != 0 {
		t.Fatal("failed")
	}
	if compareVersion("1.0", "1.0.0") != 0 {
		t.Fatal("failed")
	}
	if compareVersion("0.1", "1.1") != -1 {
		t.Fatal("failed")
	}
}
