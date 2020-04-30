package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{4, 1, 4, 6}, Answer: []int{1, 6}},
		{Input: []int{1, 2, 10, 4, 1, 4, 3, 3}, Answer: []int{2, 10}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		ans := singleNumbers(input.([]int))
		if ans[0] > ans[1] {
			ans = []int{ans[1], ans[0]}
		}
		return ans
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func singleNumbers(nums []int) []int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	dig := 1
	for dig&ans == 0 {
		dig <<= 1
	}
	a := 0
	b := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&dig > 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}
