package main

import (
	"fmt"
	"github.com/c0dehvb/leetcode/pkg/test"
	"strconv"
)

func main() {
	testCases := []test.TestCase{
		{Input: []interface{}{"12345", 2}, Answer: "45123"},
		{Input: []interface{}{"012", 4}, Answer: "201"},
		{Input: []interface{}{"012", 3}, Answer: "012"},
		{Input: []interface{}{"1", 3}, Answer: "1"},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		s := (input.([]interface{}))[0].(string)
		k := (input.([]interface{}))[1].(int)
		t, _ := strconv.Atoi(s[len(s)-1:])
		head := &ListNode{t, nil}

		for i := len(s) - 2; i >= 0; i-- {
			val, _ := strconv.Atoi(s[i : i+1])
			head = &ListNode{val, head}
		}

		head = rotateRight(head, k)

		ans := ""
		for head != nil {
			ans += fmt.Sprintf("%d", head.Val)
			head = head.Next
		}
		return ans
	})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//----------------------------------------下面才是解题代码----------------------------------------

func Len(head *ListNode) (int, *ListNode) {
	c := 0
	tail := head
	for head != nil {
		c++
		tail = head
		head = head.Next
	}
	return c, tail
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	mod, tail := Len(head)
	k %= mod
	if k == 0 {
		return head
	}

	newHead := head
	var preHead = head
	for i := 0; i < mod-k; i++ {
		preHead = newHead
		newHead = newHead.Next
	}
	preHead.Next = nil
	tail.Next = head
	return newHead
}
