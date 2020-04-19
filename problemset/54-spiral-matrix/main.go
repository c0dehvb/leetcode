package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, Answer: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{Input: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}, Answer: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{Input: [][]int{
			{3},
			{2},
		}, Answer: []int{3, 2}},
		{Input: [][]int{
			{3, 2},
		}, Answer: []int{3, 2}},
		{Input: [][]int{}, Answer: []int{}},
		{Input: [][]int{
			{},
			{},
			{},
		}, Answer: []int{}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return spiralOrder(input.([][]int))
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

func canMove(p Point2D, width, height int, visited [][]bool) bool {
	return p.X >= 0 && p.Y >= 0 &&
		p.X < width && p.Y < height &&
		!visited[p.Y][p.X]
}

func spiralOrder(matrix [][]int) []int {
	height := len(matrix)
	if height == 0 {
		return []int{}
	}
	width := len(matrix[0])
	if width == 0 {
		return []int{}
	}
	ans := []int{matrix[0][0]}
	visited := make([][]bool, height) // 记录走过的格子
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}
	// 行走方向按顺时针转动
	walkDirect := []Vector2D{
		{1, 0},  // 右
		{0, 1},  // 下
		{-1, 0}, // 左
		{0, -1}, // 上
	}
	currentPos := Point2D{0, 0}
	visited[0][0] = true

	for {
		moved := false
		for dir := 0; dir < len(walkDirect); {
			nextPos := move(currentPos, walkDirect[dir])
			if canMove(nextPos, width, height, visited) {
				ans = append(ans, matrix[nextPos.Y][nextPos.X])
				visited[nextPos.Y][nextPos.X] = true
				currentPos = nextPos
				moved = true
			} else {
				dir++
			}
		}
		if !moved {
			break
		}
	}
	return ans
}
