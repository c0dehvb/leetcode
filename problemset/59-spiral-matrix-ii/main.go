package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: 1, Answer: [][]int{
			{1},
		}},
		{Input: 2, Answer: [][]int{
			{1, 2},
			{4, 3},
		}},
		{Input: 3, Answer: [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		}},
		{Input: 4, Answer: [][]int{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 16, 15, 6},
			{10, 9, 8, 7},
		}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return generateMatrix(input.(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

type Point2D struct {
	X int
	Y int
}

type Vector2D Point2D

func move(p Point2D, v Vector2D) Point2D {
	p.X += v.X
	p.Y += v.Y
	return p
}

func canMove(p Point2D, width, height int, visited [][]int) bool {
	return p.X >= 0 && p.Y >= 0 &&
		p.X < width && p.Y < height &&
		visited[p.Y][p.X] == 0
}

func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	// 初始化矩阵
	ans := make([][]int, n)
	for i, _ := range ans {
		ans[i] = make([]int, n)
	}

	// 行走方向按顺时针转动
	walkDirect := []Vector2D{
		{1, 0},  // 右
		{0, 1},  // 下
		{-1, 0}, // 左
		{0, -1}, // 上
	}
	count := 1 // 记录已经写了多少个数
	ans[0][0] = 1
	curPos := Point2D{0, 0}
	dir := 0 // 记录当前方向

	for count < n*n {
		nextPos := move(curPos, walkDirect[dir])
		if canMove(nextPos, n, n, ans) {
			count++
			curPos = nextPos
			ans[nextPos.Y][nextPos.X] = count
		} else {
			dir = (dir + 1) % 4
		}
	}
	return ans
}
