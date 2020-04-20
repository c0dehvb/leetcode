package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
	"math"
)

func main() {
	testCases := []test.TestCase{
		{Input: 4, Answer: 2},
		{Input: 8, Answer: 2},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return mySqrt(input.(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}
