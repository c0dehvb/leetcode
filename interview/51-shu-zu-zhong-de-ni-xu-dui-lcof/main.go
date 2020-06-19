package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{7, 5, 6, 4}, Answer: 5},
	}

	n := 50000
	input := make([]int, n)
	for i := n; i > 0; i-- {
		input[n-i] = i
	}
	testCases = append(testCases, test.TestCase{input, 25000 * 49999})

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return reversePairs(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type BstTree struct {
	Root *TreeNode
}

func (bt *BstTree) Insert(node *TreeNode) {

}

func (bt *BstTree) Delete(node *TreeNode) {

}

func reversePairs(nums []int) int {
	var ans int

	return ans
}
