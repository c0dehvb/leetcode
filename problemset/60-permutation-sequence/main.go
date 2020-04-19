package main

import (
	"fmt"
	"github.com/c0dehvb/leetcode/pkg/test"
	"sort"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{3, 1}, Answer: "123"},
		{Input: []int{3, 2}, Answer: "132"},
		{Input: []int{3, 3}, Answer: "213"},
		{Input: []int{3, 6}, Answer: "321"},
		{Input: []int{4, 9}, Answer: "2314"},
		{Input: []int{4, 24}, Answer: "4321"},
		{Input: []int{1, 1}, Answer: "1"},
		{Input: []int{2, 1}, Answer: "12"},
		{Input: []int{2, 2}, Answer: "21"},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return getPermutation(input.([]int)[0], input.([]int)[1])
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

var fac = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func Swap(b []byte, i, j int) {
	b[i] ^= b[j]
	b[j] ^= b[i]
	b[i] ^= b[j]
}

func getPermutation(n int, k int) string {
	ans := make([]byte, n)
	for i := 1; i <= n; i++ {
		ans[i-1] = []byte(fmt.Sprintf("%d", i))[0]
	}
	k--
	for k > 0 {
	Loop:
		for i := len(fac) - 1; i > 0; i-- {
			if fac[i] <= k {
				for j := 9; j > 0; j-- {
					if fac[i]*j <= k {
						Swap(ans, n-i-1, n-i-1+j)
						s := n - i
						sort.Slice(ans[s:], func(i, j int) bool {
							return ans[s+i] < ans[s+j]
						})
						k -= fac[i] * j
						break Loop
					}
				}
			}
		}
	}
	return string(ans)
}
