package main

import (
	"fmt"
	"github.com/c0dehvb/leetcode/pkg/test"
	"strings"
)

func main() {
	testCases := []test.TestCase{
		{Input: "/home/", Answer: "/home"},
		{Input: "/../", Answer: "/"},
		{Input: "/home//foo/", Answer: "/home/foo"},
		{Input: "/a/./b/../../c/", Answer: "/c"},
		{Input: "/a/../../b/../c//.//", Answer: "/c"},
		{Input: "/a//b////c/d//././/..", Answer: "/a/b/c"},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return simplifyPath(input.(string))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

var (
	ErrCollectionFull  = fmt.Errorf("collection is full")
	ErrCollectionEmpty = fmt.Errorf("collection is empty")
)

type Stack struct {
	elems []string
	top   int
}

func NewStack(cap int) *Stack {
	return &Stack{make([]string, cap), 0}
}

func (s *Stack) Push(elem string) error {
	if s.top == len(s.elems) {
		return ErrCollectionFull
	}
	s.elems[s.top] = elem
	s.top++
	return nil
}

func (s *Stack) Pop() (string, error) {
	if s.top == 0 {
		return "", ErrCollectionEmpty
	}
	s.top--
	elem := s.elems[s.top]
	return elem, nil
}

func (s *Stack) Size() int {
	return s.top
}

func simplifyPath(path string) string {
	var name string
	stack := NewStack(len(path))
	idx := 0
	for idx < len(path) {
		sepIdx := strings.Index(path[idx:], "/")
		if sepIdx != -1 {
			if sepIdx == 0 {
				idx++
				continue
			}
			name = path[idx : idx+sepIdx]
			idx += sepIdx + 1
			dealName(stack, name)
		} else {
			name := path[idx:]
			dealName(stack, name)
			break
		}
	}

	var ans string
	if stack.Size() > 0 {
		for stack.Size() > 0 {
			elem, _ := stack.Pop()
			ans = "/" + elem + ans
		}
	} else {
		ans = "/"
	}
	return ans
}

func dealName(stack *Stack, name string) *Stack {
	// 处理'/'前缀
	if name[0] == '/' && len(name) > 1 {
		name = name[1:]
	}
	if name == ".." {
		if stack.Size() > 0 {
			_, _ = stack.Pop()
		}
	} else if name != "." {
		_ = stack.Push(name)
	}
	return stack
}
