package list

import "testing"

var (
	testLists1    = [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	testLists1Res = []int{1, 1, 2, 3, 4, 4, 5, 6}
)

func TestMergeKLists(t *testing.T) {
	lists := makeLists(testLists1)
	resLists := getList(mergeKLists(lists))
	if !compareArrayOfInt(resLists, testLists1Res) {
		t.Fatal("failed")
	}
}
