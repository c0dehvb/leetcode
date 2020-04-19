package main

import (
	"fmt"
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []interface{}{2.1, 3}, Answer: 9.261},
		{Input: []interface{}{2.0, -2}, Answer: 0.25},
	}

	test.Check(testCases,
		func(input interface{}) (result interface{}) {
			in := input.([]interface{})
			return myPow(in[0].(float64), in[1].(int))
		}, func(answer, result interface{}) bool {
			ans := answer.(float64)
			res := result.(float64)
			return fmt.Sprintf("%.f", ans) == fmt.Sprintf("%.f", res)
		})
}

//----------------------------------------下面才是解题代码----------------------------------------

// 快速幂
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	ans := float64(1)
	for n > 0 {
		if (n & 1) != 0 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	return ans
}
