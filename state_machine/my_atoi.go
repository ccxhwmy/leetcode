package state_machine

import "unicode"

const INT_MAX = int32(^uint32(0) >> 1)
const INT_MIN = ^INT_MAX

var (
	table = map[string][]string{
		"start":     []string{"start", "signed", "in_number", "end"},
		"signed":    []string{"end", "end", "in_number", "end"},
		"in_number": []string{"end", "end", "in_number", "end"},
		"end":       []string{"end", "end", "end", "end"},
	}
)

type autoMaton struct {
	sign int
	ans  int
	state string
}

func newAutoMaton() *autoMaton {
	return &autoMaton{
		sign:  1,
		ans:   0,
		state: "start",
	}
}

func (this *autoMaton) getCol(c rune) int {
	if unicode.IsSpace(c) {
		return 0
	}
	if c == '+' || c == '-' {
		return 1
	}
	if unicode.IsDigit(c) {
		return 2
	}
	return 3
}

func (this *autoMaton) get(c rune) {

	state := table[this.state][this.getCol(c)]
	this.state = state
	if state == "in_number" {
		this.ans = 10 * this.ans + int(c - '0')
		if this.sign == 1 {
			this.ans = min(this.ans, int(INT_MAX))
		} else {
			this.ans = min(this.ans, int(INT_MIN) * -1)
		}
	} else if (state == "signed") {
		if c == '+' {
			this.sign = 1
		} else {
			this.sign = -1
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

//https://leetcode.com/problems/string-to-integer-atoi

func myAtoi(s string) int {
	automaton := newAutoMaton()
	for _, c := range s {
		automaton.get(c)
	}
	return automaton.sign * automaton.ans
}

