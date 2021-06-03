package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{1, 2, 3, 3, 4, 4, 5}, Answer: []int{1, 2, 5}},
		{Input: []int{1, 1, 1, 2, 3}, Answer: []int{2, 3}},
		{Input: []int{}, Answer: []int{}},
		{Input: []int{1}, Answer: []int{1}},
		{Input: []int{1, 1, 1}, Answer: []int{}},
		{Input: []int{1, 1, 1, 2}, Answer: []int{2}},
		{Input: []int{0, 1, 1, 1}, Answer: []int{0}},
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
	var p *ListNode
	var pp **ListNode
	pp = &head
	for (*pp) != nil && (*pp).Next != nil {
		if (*pp).Val == (*pp).Next.Val {
			p = *pp
			val := p.Val
			for p.Next != nil && p.Val == val {
				p = p.Next
			}
			if p.Val == val {
				*pp = nil
			} else {
				*pp = p
			}
		} else {
			pp = &((*pp).Next)
		}
	}
	return head
}
