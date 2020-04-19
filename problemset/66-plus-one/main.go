package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{1, 2, 3}, Answer: []int{1, 2, 4}},
		{Input: []int{4, 3, 2, 1}, Answer: []int{4, 3, 2, 2}},
		{Input: []int{9, 9, 9}, Answer: []int{1, 0, 0, 0}},
		{Input: []int{0}, Answer: []int{1}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return plusOne(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func plusOne(digits []int) []int {
	curIdx := len(digits) - 1
	digits[curIdx] += 1

	for curIdx >= 0 && digits[curIdx] > 9 {
		digits[curIdx] %= 10
		if curIdx == 0 {
			digits = append([]int{1}, digits...)
		} else {
			digits[curIdx-1] += 1
		}
		curIdx--
	}
	return digits
}
