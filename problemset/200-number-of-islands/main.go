package main

import (
	"fmt"
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, Answer: 1},
		{Input: [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, Answer: 3},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return numIslands(input.([][]byte))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

type Point2D struct {
	X, Y int
}

type Vector2D Point2D

func (p Point2D) move(v Vector2D) Point2D {
	p.X += v.X
	p.Y += v.Y
	return p
}

type Queue struct {
	elems []Point2D
	l, r  int
}

var (
	ErrQueueFull  = fmt.Errorf("queue is full")
	ErrQueueEmpty = fmt.Errorf("queue is empty")
)

func NewQueue(cap int) *Queue {
	return &Queue{
		elems: make([]Point2D, cap),
		l:     0,
		r:     0,
	}
}

func (q *Queue) Push(d Point2D) error {
	if q.r < len(q.elems) {
		q.elems[q.r] = d
		q.r++
		return nil
	}
	return ErrQueueFull
}

func (q *Queue) Pop() (Point2D, error) {
	p := Point2D{}
	if q.Size() == 0 {
		return p, ErrQueueEmpty
	} else {
		p = q.elems[q.l]
		q.l++
		return p, nil
	}
}

func (q *Queue) Size() int {
	return q.r - q.l
}

func numIslands(grid [][]byte) int {
	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])
	if width == 0 {
		return 0
	}
	islandSet := make([][]int, height)
	for i := 0; i < height; i++ {
		islandSet[i] = make([]int, width)
		for j := 0; j < width; j++ {
			islandSet[i][j] = width*height + 1
		}
	}

	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '1' && islandSet[y][x] == width*height+1 {
				count++
				islandSet = bfsUnionSet(grid, islandSet, Point2D{x, y}, count)
			}
		}
	}
	return count
}

// BFS并查集
func bfsUnionSet(grid [][]byte, islandSet [][]int, start Point2D, setNo int) [][]int {
	height := len(grid)
	width := len(grid[0])
	queue := NewQueue(width * height)
	_ = queue.Push(start)
	islandSet[start.Y][start.X] = setNo

	// 上下左右
	direct := []Vector2D{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	for queue.Size() != 0 {
		curPos, _ := queue.Pop()
		// 扫描相邻4个格子
		for _, d := range direct {
			nextPos := curPos.move(d)
			if nextPos.X >= 0 && nextPos.X < width &&
				nextPos.Y >= 0 && nextPos.Y < height &&
				grid[nextPos.Y][nextPos.X] == '1' &&
				islandSet[nextPos.Y][nextPos.X] > islandSet[curPos.Y][curPos.X] {
				islandSet[nextPos.Y][nextPos.X] = islandSet[curPos.Y][curPos.X]
				_ = queue.Push(nextPos)
			}
		}
	}
	return islandSet
}
