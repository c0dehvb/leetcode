package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
	"regexp"
	"strings"
)

func main() {
	testCases := []test.TestCase{
		{Input: []string{"aa", "a"}, Answer: false},
		{Input: []string{"aa", "*"}, Answer: true},
		{Input: []string{"cb", "?a"}, Answer: false},
		{Input: []string{"adceb", "*a*b"}, Answer: true},
		{Input: []string{"acdcb", "a*c?b"}, Answer: false},
		{Input: []string{"acdcb", "a****c????b"}, Answer: false},
		{Input: []string{"acdcb", "a****c??b"}, Answer: true},
		{Input: []string{"acdcbb", "*a****c?*?b?"}, Answer: true},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		in := input.([]string)
		return isMatch(in[0], in[1])
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func isMatch(s string, p string) bool {
	p = strings.ReplaceAll(p, "?", ".")
	p = strings.ReplaceAll(p, "*", ".*")
	p = "^" + p + "$"
	ans, _ := regexp.Match(p, []byte(s))
	return ans
}