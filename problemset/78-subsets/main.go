package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{1, 2, 3}, Answer: [][]int{
			{},
			{1},
			{1, 2},
			{1, 2, 3},
			{1, 3},
			{2},
			{2, 3},
			{3},
		}},
		{Input: []int{}, Answer: [][]int{
			{},
		}},
		{Input: []int{1}, Answer: [][]int{
			{},
			{1},
		}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return subsets(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func subsets(nums []int) [][]int {
	var ans [][]int
	ans = append(ans, []int{})
	// 从左到右遍历
	var stack []int
	stack = append(stack, 0)
	returnUp := func() {
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			stack[len(stack)-1]++
		}
	}
	for len(stack) > 0 {
		aSet, success := buildResult(nums, stack)
		if !success {
			returnUp()
			continue
		}
		ans = append(ans, aSet)
		if stack[len(stack)-1] < len(nums)-1 {
			stack = append(stack, stack[len(stack)-1]+1)
		} else {
			returnUp()
		}
	}
	return ans
}

func buildResult(nums, indexes []int) ([]int, bool) {
	ans := make([]int, len(indexes))
	for i, idx := range indexes {
		if idx < len(nums) {
			ans[i] = nums[idx]
		} else {
			return nil, false
		}
	}
	return ans, true
}
