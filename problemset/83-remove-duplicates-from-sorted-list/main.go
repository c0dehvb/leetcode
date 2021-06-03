package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{1, 1, 2}, Answer: []int{1, 2}},
		{Input: []int{1, 1, 2, 3, 3}, Answer: []int{1, 2, 3}},
		{Input: []int{1, 1, 1, 1}, Answer: []int{1}},
		{Input: []int{}, Answer: []int{}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		in := buildListNode(input.([]int))
		return deleteDuplicates(in).ToArray()
	})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{}
	var p = head
	for i := 0; i < len(arr); i++ {
		p.Val = arr[i]
		if i != len(arr)-1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	return head
}

func (l *ListNode) ToArray() []int {
	p := l
	var ans []int
	for p != nil {
		ans = append(ans, p.Val)
		p = p.Next
	}
	return ans
}

//----------------------------------------下面才是解题代码----------------------------------------

// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序排列
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			for p.Next != nil && p.Next.Val == p.Val {
				p.Next = p.Next.Next
			}
		}
		p = p.Next
	}
	return head
}
