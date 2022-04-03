package dynamic

//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	min := prices[0]
	for _, price := range prices {
		if min > price {
			min = price
		}
		if res < price - min {
			res = price - min
		}
	}
	return res
}
