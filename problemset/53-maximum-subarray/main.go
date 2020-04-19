package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, Answer: 6},
		{Input: []int{-1, -2, -3}, Answer: -1},
		{Input: []int{1,0,0,0,0,0,1}, Answer: 2},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return maxSubArray(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func maxSubArray(nums []int) int {
	ans := nums[0]
	maxSub := make([]int, len(nums))
	maxSub[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		maxSub[i] = maxSub[i-1] + nums[i]
		if maxSub[i-1] < 0 {
			maxSub[i] = nums[i]
		}
		if ans < maxSub[i] {
			ans = maxSub[i]
		}
	}
	return ans
}