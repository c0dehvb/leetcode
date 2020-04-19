package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
	"sort"
)

func main() {
	testCases := []test.TestCase{
		{Input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, Answer: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{Input: [][]int{{1, 4}, {4, 5}}, Answer: [][]int{{1, 5}}},
		{Input: [][]int{}, Answer: [][]int{}},
		{Input: [][]int{{1, 2}, {1, 3}, {1, 4}}, Answer: [][]int{{1, 4}}},
		{Input: [][]int{{1, 2}, {1, 2}, {1, 2}}, Answer: [][]int{{1, 2}}},
		{Input: [][]int{{1, 2}, {2, 3}, {3, 4}}, Answer: [][]int{{1, 4}}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return merge(input.([][]int))
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

// 先排序，然后一次遍历合并到新数组
// 时间复杂度: O(NlogN)
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	segSlice := &SegmentSlice{list: []Segment{}}
	for _, s := range intervals {
		segSlice.list = append(segSlice.list, Segment{s[0], s[1]})
	}
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
