package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{2, 3, 1, 1, 4}, Answer: true},
		{Input: []int{3, 2, 1, 0, 4}, Answer: false},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return canJump(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

// 0-1背包
func canJump(nums []int) bool {
	visited := make([]bool, len(nums))
	visited[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums) && j-i <= nums[i]; j++ {
			if visited[j] {
				visited[i] = true
				break
			}
		}
	}
	return visited[0]
}
