package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, Answer: 6},
		{Input: []int{0, 5, 4, 3, 2, 1, 0, 5, 0}, Answer: 15},
		{Input: []int{5, 4, 3, 2, 1, 0, 5, 0}, Answer: 15},
		{Input: []int{0, 5, 4, 3, 2, 1, 0, 5}, Answer: 15},
		{Input: []int{5, 4, 3, 2, 1, 0, 5}, Answer: 15},
		{Input: []int{0, 5, 4, 3, 2, 1, 0}, Answer: 0},
		{Input: []int{5, 2, 1, 2, 1, 5}, Answer: 14},
		{Input: []int{5, 0, 1, 2, 3, 4, 5}, Answer: 15},
		{Input: []int{5, 4, 3, 2, 1, 2, 3, 4, 5}, Answer: 16},
		{Input: []int{5, 0, 0, 0, 5}, Answer: 15},
		{Input: []int{5, 0, 0, 0, 5, 0, 0, 0, 5}, Answer: 30},
		{Input: []int{5, 0, 0, 0, 5, 0, 0, 0, 0}, Answer: 15},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return trap(input.([]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

type Grid struct {
	Height int
	Pile   int
}

func trap(height []int) int {
	result := 0

	// 构建压缩图
	grids := []Grid{}
	for i := 0; i < len(height); i++ {
		g := Grid{Height: height[i], Pile: 1}
		grids = append(grids, g)
	}

	hasGet := true
	for hasGet {
		hasGet = false
		grids = compress(grids)

		for i := 1; i < len(grids)-1; i++ {
			if grids[i].Height < grids[i-1].Height && grids[i].Height < grids[i+1].Height {
				var l, r int
				for l = i - 1; l > 0 && grids[l-1].Height > grids[l].Height; l-- {
				}
				for r = i + 1; r < len(grids)-1 && grids[r].Height < grids[r+1].Height; r++ {
				}
				get := 0
				if grids[r].Height < grids[l].Height {
					for j := r - 1; j >= l && grids[r].Height > grids[j].Height; j-- {
						get += (grids[r].Height - grids[j].Height) * grids[j].Pile
						grids[j].Height = grids[r].Height
					}
				} else {
					for j := l + 1; j <= r && grids[l].Height > grids[j].Height; j++ {
						get += (grids[l].Height - grids[j].Height) * grids[j].Pile
						grids[j].Height = grids[l].Height
					}
				}
				hasGet = true
				result += get
				break
			}
		}
	}
	return result
}

// 构建压缩图
func compress(in []Grid) []Grid {
	var i, j int
	out := []Grid{}
	for i = 0; i < len(in); i = j {
		if in[i].Pile == 0 {
			continue
		}
		g := Grid{
			Height: in[i].Height,
			Pile:   in[i].Pile,
		}
		for j = i + 1; j < len(in) && in[j].Height == in[i].Height; j++ {
			g.Pile += in[j].Pile
		}
		out = append(out, g)
	}
	return out
}
