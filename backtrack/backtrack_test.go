package backtrack

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	fmt.Println(permuteWithRecursion([]int{1, 2, 3}))
	fmt.Println(permute([]int{1, 2, 3}))
}
