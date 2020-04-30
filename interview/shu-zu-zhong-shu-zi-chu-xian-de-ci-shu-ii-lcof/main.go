package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{3, 4, 3, 3}, Answer: 4},
		{Input: []int{9, 1, 7, 9, 7, 9, 7}, Answer: 1},
		{Input: []int{1}, Answer: 1},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return singleNumber(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func singleNumber(nums []int) int {
	byteCount := make([]int, 32)
	for _, num := range nums {
		for dig := 0; num > 0; num >>= 1 {
			if num&1 == 1 {
				byteCount[dig] = (byteCount[dig] + 1) % 3
			}
			dig++
		}
	}
	ans := 0
	for i := 31; i >= 0; i-- {
		ans = (ans << 1) | byteCount[i]
	}
	return ans
}
