package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: 5, Answer: 2},
		{Input: 10, Answer: 4},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return waysToChange(input.(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

const MOD = 1000000007

var coins = []int{1, 5, 10, 25}

func waysToChange(n int) int {
	ans := make([]int, n+1)
	ans[0] = 1

	for j := 0; j < len(coins); j++ {
		for i := 1; i <= n; i++ {
			if i >= coins[j] {
				ans[i] = (ans[i] + ans[i-coins[j]]) % MOD
			}
		}
	}
	return ans[n]
}
