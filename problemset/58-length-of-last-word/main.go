package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: "Hello world", Answer: 5},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return lengthOfLastWord(input.(string))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && ans > 0 {
			return ans
		} else if s[i] != ' ' {
			ans++
		}
	}
	return ans
}
