package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []interface{}{[]int{1, 1, 2, 1, 1}, 3}, Answer: 2},
		{Input: []interface{}{[]int{2, 4, 6}, 1}, Answer: 0},
		{Input: []interface{}{[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2}, Answer: 16},
		{Input: []interface{}{[]int{1}, 2}, Answer: 0},
		{Input: []interface{}{[]int{1, 1, 1}, 2}, Answer: 2},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return numberOfSubarrays(input.([]interface{})[0].([]int), input.([]interface{})[1].(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func numberOfSubarrays(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}
	var l, r, ans int
	oddNum := 0
	for ; len(nums)-r >= k-oddNum && r < len(nums); r++ {
		if nums[r]%2 != 0 {
			oddNum++
		}
		if oddNum == k {
			t1 := 1
			t2 := 0
			for r+1 < len(nums) && nums[r+1]%2 == 0 {
				t1++
				r++
			}
			for oddNum == k {
				t2++
				if nums[l]%2 != 0 {
					oddNum--
				}
				l++
			}
			ans += t2 * t1
		}
	}
	return ans
}
