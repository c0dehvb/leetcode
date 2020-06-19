package main

import "github.com/c0dehvb/leetcode/pkg/test"

func main() {
	testCases := []test.TestCase{
		{Input: []string{"ADOBECODEBANC", "ABC"}, Answer: "BANC"},
		{Input: []string{"ABCDABCDABCDABCD", "DCBA"}, Answer: "ABCD"},
		{Input: []string{"ABCDABCDABCD", "ABCD"}, Answer: "ABCD"},
		{Input: []string{"ABCCCCCCCCCCCCCCCD", "DCBA"}, Answer: "ABCCCCCCCCCCCCCCCD"},
		{Input: []string{"DCCCCCCCCCCCCCCABD", "DCBA"}, Answer: "CABD"},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return minWindow(input.([]string)[0], input.([]string)[1])
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func minWindow(s string, t string) string {
	checkNum := 0
	minLen := len(s)
	ans := ""
	chars := buildCharMap(t)
	var l, r int

	for r < len(s) {
		c := s[r]
		if foundChar(c, chars) {
			chars[c] -= 1
			if chars[c] >= 0 {
				checkNum += 1
			}

			for checkNum == len(t) {
				if r-l < minLen {
					ans = s[l : r+1]
					minLen = r - l
				}

				if foundChar(s[l], chars) {
					chars[s[l]] += 1
					if chars[s[l]] > 0 {
						checkNum -= 1
					}
				}
				l++
			}
		}
		r++
	}
	return ans
}

func buildCharMap(s string) map[byte]int {
	chars := make(map[byte]int)
	for _, c := range s {
		chars[byte(c)] += 1
	}
	return chars
}

func foundChar(c byte, chars map[byte]int) bool {
	_, found := chars[c]
	return found
}
