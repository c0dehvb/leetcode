package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: 19, Answer: true},
		{Input: 1, Answer: true},
		{Input: 2, Answer: false},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return isHappy(input.(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = true
		// next num
		next := 0
		for n > 0 {
			next += (n % 10) * (n % 10)
			n /= 10
		}
		n = next
	}
	return n == 1
}
