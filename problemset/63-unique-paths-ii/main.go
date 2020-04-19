package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}, Answer: 2},
		{Input: [][]int{
			{0, 0, 0},
			{1, 1, 1},
			{0, 0, 0},
		}, Answer: 0},
		{Input: [][]int{
			{0, 0, 0},
			{0, 1, 1},
			{0, 0, 0},
		}, Answer: 1},
		{Input: [][]int{
			{1},
		}, Answer: 0},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return uniquePathsWithObstacles(input.([][]int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

// 动态规划问题
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])
	if obstacleGrid[height-1][width-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}
	ans := make([][]int, height)
	for i := 0; i < height; i++ {
		ans[i] = make([]int, width)
	}
	ans[height-1][width-1] = 1
	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			if obstacleGrid[y][x] == 0 {
				if y+1 < height {
					ans[y][x] += ans[y+1][x]
				}
				if x+1 < width {
					ans[y][x] += ans[y][x+1]
				}
			}
		}
	}
	return ans[0][0]
}
