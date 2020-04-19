package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{2, 3, 1, 1, 4}, Answer: 2},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return jump(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func jump(nums []int) int {
	steps := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		steps[i] = len(nums) + 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j && steps[j] > steps[i]+1 {
				steps[j] = steps[i] + 1
			}
		}
	}
	return steps[0]
}
