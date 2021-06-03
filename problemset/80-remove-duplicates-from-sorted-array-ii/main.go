package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{1, 1, 1, 2, 2, 3}, Answer: 5},
		{Input: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, Answer: 7},
		{Input: []int{0}, Answer: 1},
	}

	input := make([]int, 3*104)
	for i := 0; i < 104; i++ {
		input[i*3] = i
		input[i*3+1] = i
		input[i*3+2] = i
	}
	testCases = append(testCases, test.TestCase{Input: input, Answer: 104 * 2})

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return removeDuplicates(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

// 1 <= nums.length <= 3 * 104
// -104 <= nums[i] <= 104
// nums 已按升序排列
func removeDuplicates(nums []int) int {
	count := 1
	num := nums[0]
	x := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			count += 1
		} else {
			count = 1
			num = nums[i]
		}
		if count <= 2 {
			nums[x] = nums[i]
			x += 1
		}
	}
	nums = nums[:x]
	return len(nums)
}
