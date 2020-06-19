package main

import (
	"fmt"
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{4, 0}, Answer: [][]int{{}}},
		{Input: []int{4, 1}, Answer: [][]int{
			{1},
			{2},
			{3},
			{4},
		}},
		{Input: []int{4, 2}, Answer: [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		}},
		{Input: []int{4, 3}, Answer: [][]int{
			{1, 2, 3},
			{1, 2, 4},
			{1, 3, 4},
			{2, 3, 4},
		}},
		{Input: []int{4, 4}, Answer: [][]int{
			{1, 2, 3, 4},
		}},
		{Input: []int{4, 5}, Answer: [][]int{}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return combine(input.([]int)[0], input.([]int)[1])
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func combine(n int, k int) [][]int {
	if k == 0 {
		return [][]int{{}}
	} else if k > n {
		return [][]int{}
	}
	count := C(n, k)
	stack := NewStack(k)
	ans := make([][]int, 0, count)

	_ = stack.Push(1)

	deep := 1
	var preNum, nextNum int
	var err error
	for stack.Size() > 0 {
		if deep == k {
			ans = append(ans, stack.Elems())
			deep -= 1
		} else {
			if stack.Size() > deep { // 从下一层return回上一层
				nextNum, _ = stack.Pop()

				if nextNum < n { // 探下一个值
					_ = stack.Push(nextNum + 1)
					deep += 1
					continue
				} else { // 探到底了，继续往上return
					deep -= 1
					continue
				}
			} else {
				preNum, err = stack.Peek()
				if err != nil {
					nextNum = 1
				} else if preNum < n {
					nextNum = preNum + 1
					_ = stack.Push(nextNum)
					deep += 1
					continue
				} else {
					deep -= 1
					continue
				}
			}
		}
	}
	return ans
}

func C(n, m int) int {
	if m > n-m {
		m = n - m
	}
	if m == 0 {
		return 1
	}
	ans := 1
	for i := 0; i < m; i++ {
		ans = ans * (n - i) / (i + 1)
	}
	return ans
}

var (
	ErrCollectionFull  = fmt.Errorf("collection is full")
	ErrCollectionEmpty = fmt.Errorf("collection is empty")
)

type Stack struct {
	elems []int
	top   int
}

func NewStack(cap int) *Stack {
	return &Stack{make([]int, cap), 0}
}

func (s *Stack) Push(elem int) error {
	if s.top == len(s.elems) {
		return ErrCollectionFull
	}
	s.elems[s.top] = elem
	s.top++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.top == 0 {
		return 0, ErrCollectionEmpty
	}
	s.top--
	elem := s.elems[s.top]
	return elem, nil
}

func (s *Stack) Peek() (int, error) {
	if s.top == 0 {
		return 0, ErrCollectionEmpty
	}
	elem := s.elems[s.top-1]
	return elem, nil
}

func (s *Stack) Size() int {
	return s.top
}

func (s *Stack) Elems() []int {
	if s.top == 0 {
		return []int{}
	}
	arr := make([]int, s.top)
	for i := 0; i < s.top; i++ {
		arr[i] = s.elems[i]
	}
	return arr
}
