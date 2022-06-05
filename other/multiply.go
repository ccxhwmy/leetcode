package other

import (
	"strconv"
)

//https://leetcode.com/problems/multiply-strings

func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := ""
	for i := m - 1; i >= 0; i-- {
		cur := ""
		add := 0
		for j := m - 1; j > i; j-- {
			cur += "0"
		}
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			product := x * y + add
			cur = strconv.Itoa(product % 10) + cur
			add = product / 10
		}
		for ; add != 0; add /= 10 {
			cur = strconv.Itoa(add % 10) + cur
		}
		res = addStrings(res, cur)
	}
	return res
}

func addStrings(num1, num2 string) string {
	add := 0
	ans := ""
	for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0 || add != 0; i, j = i - 1, j - 1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		res := x + y + add
		ans = strconv.Itoa(res % 10) + ans
		add = res / 10
	}
	return ans
}
