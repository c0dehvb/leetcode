package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{}, Answer: []int{}},
		{Input: []int{1}, Answer: []int{1}},
		{Input: []int{1, 2, 3, 0, 5, 0, 4}, Answer: []int{1, 3, 4}},
		{Input: []int{1, 2, 3, 6, 5, 0, 4, 7, 0}, Answer: []int{1, 3, 4, 7}},
		{Input: []int{1, 2, 3, 6, 5, 0, 4, 7, 8}, Answer: []int{1, 3, 4, 8}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		in := input.([]int)
		if len(in) == 0 {
			return rightSideView(nil)
		}
		nodes := make([]*TreeNode, len(in))
		for i, _ := range in {
			if in[i] == 0 {
				continue
			}
			nodes[i] = &TreeNode{Val: in[i], Left: nil, Right: nil}
		}
		for i, _ := range in {
			if nodes[i] == nil {
				continue
			}
			l := (i+1)*2 - 1
			r := l + 1
			if l < len(in) {
				nodes[i].Left = nodes[l]
			}
			if r < len(in) {
				nodes[i].Right = nodes[r]
			}
		}
		return rightSideView(nodes[0])
	})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//----------------------------------------下面才是解题代码----------------------------------------

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	visit := make(map[int]int)
	rightSideViewInternal(root, 0, visit)
	ans := make([]int, len(visit))
	for i, v := range visit {
		ans[i] = v
	}
	return ans
}

func rightSideViewInternal(node *TreeNode, deep int, visit map[int]int) {
	if _, found := visit[deep]; !found {
		visit[deep] = node.Val
	}
	if node.Right != nil {
		rightSideViewInternal(node.Right, deep+1, visit)
	}
	if node.Left != nil {
		rightSideViewInternal(node.Left, deep+1, visit)
	}
}
