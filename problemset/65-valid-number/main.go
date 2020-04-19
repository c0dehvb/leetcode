package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
	"strings"
)

func main() {
	testCases := []test.TestCase{
		{Input: "0", Answer: true},
		{Input: "0.1", Answer: true},
		{Input: "abc", Answer: false},
		{Input: "1 a", Answer: false},
		{Input: "2e10", Answer: true},
		{Input: " -90e3   ", Answer: true},
		{Input: " 1e", Answer: false},
		{Input: "e3", Answer: false},
		{Input: " 6e-1", Answer: true},
		{Input: " 6e+1", Answer: true},
		{Input: " 99e2.5 ", Answer: false},
		{Input: "53.5e93", Answer: true},
		{Input: " --6 ", Answer: false},
		{Input: "-+3", Answer: false},
		{Input: "95a54e53", Answer: false},
		{Input: ".1", Answer: true},
		{Input: "3.", Answer: true},
		{Input: ".", Answer: false},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return isNumber(input.(string))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if eIdx := strings.Index(s, "e"); eIdx != -1 {
		return isNumberBeforeE(s[:eIdx]) && isNumberAfterE(s[eIdx+1:])
	} else {
		return isNumberBeforeE(s)
	}
}

func isNumberBeforeE(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '-' || s[0] == '+' {
		if len(s) > 1 {
			s = s[1:]
		} else {
			return false
		}
	}
	hasPoint := false
	for i, c := range s {
		if c == '.' && !hasPoint {
			hasPoint = true
			// ".1" 与 "3." 题解认为是 true
			//if i == 0 || i == len(s)-1 {
			//	return false
			//}
			// 但 "." 题解认为是 false
			if i == 0 && i == len(s)-1 {
				return false
			}
		} else if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func isNumberAfterE(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '-' || s[0] == '+' {
		if len(s) > 1 {
			s = s[1:]
		} else {
			return false
		}
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
