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
	splits := strings.Split(path, "/")
	if len(splits) == 0 {
		return "/"
	}
	stack := NewStack(len(path))
	for _, name := range splits {
		if name == ".." {
			if stack.Size() > 0 {
				_, _ = stack.Pop()
			}
		} else if name != "." && name != "" {
			_ = stack.Push(name)
		}
	}

	var ans string
	for stack.Size() > 0 {
		elem, _ := stack.Pop()
		ans = "/" + elem + ans
	}
	if ans == "" {
		ans = "/"
	}
	return ans
}
