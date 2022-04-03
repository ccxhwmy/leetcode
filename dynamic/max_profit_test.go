package dynamic

import "testing"

var (
	prices1 = []int{7, 1, 5, 3, 6, 4}
	prices2 = []int{1, 2, 3, 4, 5}
	prices3 = []int{7, 6, 4, 3, 1}
	prices4 = []int{3, 3, 5, 0, 0, 3, 1, 4}
)

func TestMaxProfit1(t *testing.T) {
	if maxProfit1(prices1) != 5 {
		t.Fatal("failed")
	}
}

func TestMaxProfit2WithDynamic(t *testing.T) {
	if maxProfit2WithDynamic(prices1) != 7 {
		t.Fatal("failed 1")
	}
	if maxProfit2WithDynamic(prices2) != 4 {
		t.Fatal("failed 2")
	}
	if maxProfit2WithDynamic(prices3) != 0 {
		t.Fatal("failed 3")
	}
}

func TestMaxProfit2WithDynamic2(t *testing.T) {
	if maxProfit2WithDynamic2(prices1) != 7 {
		t.Fatal("failed 1")
	}
	if maxProfit2WithDynamic2(prices2) != 4 {
		t.Fatal("failed 2")
	}
	if maxProfit2WithDynamic2(prices3) != 0 {
		t.Fatal("failed 3")
	}
}

func TestMaxProfit2WithGreedy(t *testing.T) {
	if maxProfit2WithGreedy(prices1) != 7 {
		t.Fatal("failed 1")
	}
	if maxProfit2WithGreedy(prices2) != 4 {
		t.Fatal("failed 2")
	}
	if maxProfit2WithGreedy(prices3) != 0 {
		t.Fatal("failed 3")
	}
}

func TestMaxProfit3(t *testing.T) {
	if maxProfit3(prices2) != 4 {
		t.Fatal("failed2")
	}
	if maxProfit3(prices3) != 0 {
		t.Fatal("failed3")
	}
	if maxProfit3(prices4) != 6 {
		t.Fatal("failed4")
	}
}
