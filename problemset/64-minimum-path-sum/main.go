package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}, Answer: 7},
		{Input: [][]int{
			{1},
		}, Answer: 1},
		{Input: [][]int{
			{1, 2},
		}, Answer: 3},
		{Input: [][]int{
			{1},
			{2},
		}, Answer: 3},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return minPathSum(input.([][]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

// 题目声明格子均为非负数，但没有给出最大值，所以假设-1为无限大
const INF = -1

func min(a, b int) int {
	if a == INF {
		return b
	} else if b == INF {
		return a
	} else if a < b {
		return a
	} else {
		return b
	}
}

func minPathSum(grid [][]int) int {
	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])
	if width == 0 {
		return 0
	}

	ans := make([][]int, height)
	for i := 0; i < height; i++ {
		ans[i] = make([]int, width)
		for j := 0; j < width; j++ {
			ans[i][j] = INF
		}
	}

	ans[height-1][width-1] = grid[height-1][width-1]
	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			if x-1 >= 0 {
				ans[y][x-1] = min(ans[y][x-1], ans[y][x]+grid[y][x-1])
			}
			if y-1 >= 0 {
				ans[y-1][x] = min(ans[y-1][x], ans[y][x]+grid[y-1][x])
			}
		}
	}
	return ans[0][0]
}
