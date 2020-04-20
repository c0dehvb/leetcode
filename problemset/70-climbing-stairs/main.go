package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: 1, Answer: 1},
		{Input: 2, Answer: 2},
		{Input: 3, Answer: 3},
		{Input: 4, Answer: 5},
		{Input: 5, Answer: 8},
		{Input: 6, Answer: 13},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return climbStairs(input.(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

// 斐波那契数列, O(n)
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	var fn, fn_1, fn_2 int
	fn_1 = 2
	fn_2 = 1
	for i := 3; i <= n; i++ {
		fn = fn_1 + fn_2
		fn_2 = fn_1
		fn_1 = fn
	}
	return fn
}
