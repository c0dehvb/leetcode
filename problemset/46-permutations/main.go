package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{1, 2, 3}, Answer: [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return permute(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		sub := append([]int{}, nums[:i]...)
		sub = append(sub, nums[i+1:]...)
		subAns := permute(sub)
		for _, s := range subAns {
			ans = append(ans, append([]int{nums[i]}, s...))
		}
	}
	return ans
}
