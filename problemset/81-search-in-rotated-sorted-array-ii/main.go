package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []interface{}{[]int{2, 5, 6, 0, 0, 1, 2}, 0}, Answer: true},
		{Input: []interface{}{[]int{2, 5, 6, 0, 0, 1, 2}, 3}, Answer: false},
		{Input: []interface{}{[]int{1, 2, 2, 5, 6, 0, 0}, 0}, Answer: true},
		{Input: []interface{}{[]int{1, 2, 2, 5, 6, 0, 0}, 3}, Answer: false},
		{Input: []interface{}{[]int{6, 0, 0, 1, 2, 2, 5}, 0}, Answer: true},
		{Input: []interface{}{[]int{6, 0, 0, 1, 2, 2, 5}, 3}, Answer: false},
		{Input: []interface{}{[]int{2, 5, 6, 0, 1, 2}, 0}, Answer: true},
		{Input: []interface{}{[]int{2, 5, 6, 0, 1, 2}, 3}, Answer: false},
		{Input: []interface{}{[]int{1, 2, 2, 5, 6, 0}, 0}, Answer: true},
		{Input: []interface{}{[]int{1, 2, 2, 5, 6, 0}, 3}, Answer: false},
		{Input: []interface{}{[]int{6, 0, 1, 2, 2, 5}, 0}, Answer: true},
		{Input: []interface{}{[]int{6, 0, 1, 2, 2, 5}, 3}, Answer: false},
		{Input: []interface{}{[]int{6}, 6}, Answer: true},
		{Input: []interface{}{[]int{6}, 3}, Answer: false},
		{Input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2}, 4}, Answer: true},
		{Input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2}, 6}, Answer: true},
		{Input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2}, 7}, Answer: true},
		{Input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, Answer: true},
		{Input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2}, 2}, Answer: true},
		{Input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 0}, Answer: true},
		{Input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 3}, Answer: true},
		{Input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 4}, Answer: true},
		{Input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 7}, Answer: true},
		{Input: []interface{}{[]int{4, 4, 5, 5, 6, 6, 7, 7, 0, 0, 1, 1, 2, 2, 3, 3}, 0}, Answer: true},
		{Input: []interface{}{[]int{4, 4, 5, 5, 6, 6, 7, 7, 0, 0, 1, 1, 2, 2, 3, 3}, 3}, Answer: true},
		{Input: []interface{}{[]int{4, 4, 5, 5, 6, 6, 7, 7, 0, 0, 1, 1, 2, 2, 3, 3}, 4}, Answer: true},
		{Input: []interface{}{[]int{4, 4, 5, 5, 6, 6, 7, 7, 0, 0, 1, 1, 2, 2, 3, 3}, 7}, Answer: true},
		{Input: []interface{}{[]int{4, 4, 5, 5, 6, 6, 7, 7, 0, 0, 1, 1, 2, 2, 3, 3}, 8}, Answer: false},
		{Input: []interface{}{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2}, Answer: true},
		{Input: []interface{}{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 3}, Answer: false},
		{Input: []interface{}{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2}, Answer: true},
		{Input: []interface{}{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 3}, Answer: false},
		{Input: []interface{}{[]int{1, 1, 1, 1, 1}, 2}, Answer: false},
		{Input: []interface{}{[]int{1, 1, 1, 1, 1}, 1}, Answer: true},
		{Input: []interface{}{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 13}, Answer: true},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return search(input.([]interface{})[0].([]int), input.([]interface{})[1].(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

// 1 <= nums.length <= 5000
//-104 <= nums[i] <= 104
// 题目数据保证 nums 在预先未知的某个下标上进行了旋转
// 有很多想法，到最后发现数组中仅有一个值不同时，你还是得一个个遍历，时间复杂度O(N)
func search(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return true
		}
	}
	return false
}
