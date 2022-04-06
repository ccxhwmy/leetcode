package slide_win

import "math"

//https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	resL, resR := -1, -1
	winLen := math.MaxInt32
	for l, r := 0, 0; r < len(s); r++ {
		if r < len(s) && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if (r - l + 1 < winLen) {
				winLen = r - l + 1
				resL, resR = l, l + winLen
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if resL == -1 {
		return ""
	}
	return s[resL:resR]
}
