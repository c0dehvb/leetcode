package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
	"sort"
)

type Input struct {
	intervals   [][]int
	newInterval []int
}

func main() {
	testCases := []test.TestCase{
		{Input: Input{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, Answer: [][]int{{1, 5}, {6, 9}}},
		{Input: Input{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}}, Answer: [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{Input: Input{[][]int{}, []int{4, 8}}, Answer: [][]int{{4, 8}}},
		{Input: Input{[][]int{}, []int{4, 8}}, Answer: [][]int{{4, 8}}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return insert(input.(Input).intervals, input.(Input).newInterval)
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

type Segment struct {
	L, R int
}

type SegmentSlice struct {
	list []Segment
}

// 实现sort.Interface接口
func (s *SegmentSlice) Len() int {
	return len(s.list)
}

// 实现sort.Interface接口
func (s *SegmentSlice) Less(i, j int) bool {
	if s.list[i].L != s.list[j].L {
		return s.list[i].L < s.list[j].L
	} else {
		return s.list[i].R < s.list[j].R
	}
}

// 实现sort.Interface接口
func (s *SegmentSlice) Swap(i, j int) {
	tmp := s.list[i]
	s.list[i] = s.list[j]
	s.list[j] = tmp
}

// 使用56题的代码简单做了点修改
// 时间复杂度: O(NlogN)
func insert(intervals [][]int, newInterval []int) [][]int {
	segSlice := &SegmentSlice{list: []Segment{}}
	for _, s := range intervals {
		segSlice.list = append(segSlice.list, Segment{s[0], s[1]})
	}
	segSlice.list = append(segSlice.list, Segment{newInterval[0], newInterval[1]})
	sort.Sort(segSlice)

	newSeg := [][]int{}
	curSeg := segSlice.list[0]

	for _, s := range segSlice.list {
		if curSeg.R < s.L {
			newSeg = append(newSeg, []int{curSeg.L, curSeg.R})
			curSeg = s
		} else if curSeg.R < s.R {
			curSeg.R = s.R
		}
	}
	newSeg = append(newSeg, []int{curSeg.L, curSeg.R})
	return newSeg
}
