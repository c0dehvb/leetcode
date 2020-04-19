package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{3, 2}, Answer: 3},
		{Input: []int{7, 3}, Answer: 28},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return uniquePaths(input.([]int)[0], input.([]int)[1])
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func C(m, n int) int {
	if n > m-n {
		n = m - n
	}
	ans := 1
	for i := 1; i <= n; i++ {
		ans = ans * (m - i + 1) / i
	}
	return ans
}

// 有重复的排列问题，主要小心算C(m,n)不要数值溢出
func uniquePaths(m int, n int) int {
	return C(m-1+n-1, n-1)
}
