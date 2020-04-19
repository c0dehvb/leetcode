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
		}, Answer: [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}},
		{Input: [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}, Answer: [][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		}},
		{Input: [][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		}, Answer: [][]int{
			{21, 16, 11, 6, 1},
			{22, 17, 12, 7, 2},
			{23, 18, 13, 8, 3},
			{24, 19, 14, 9, 4},
			{25, 20, 15, 10, 5},
		}},
	}

	test.SimpleCheck(testCases,
		func(input interface{}) (result interface{}) {
			in := input.([][]int)
			rotate(in)
			return in
		})
}

//----------------------------------------下面才是解题代码----------------------------------------

type Point2D struct {
	X int
	Y int
}

// 题目要求用原矩阵来旋转图像，所以需要分开每一轮交换4个点的坐标
func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < (len(matrix[i])+1)/2; j++ {
			p := Point2D{X: j, Y: i}
			pSet := calcRotateSet(p, len(matrix[i]))
			tmp := matrix[pSet[3].Y][pSet[3].X]
			for k := 3; k > 0; k-- {
				matrix[pSet[k].Y][pSet[k].X] = matrix[pSet[k-1].Y][pSet[k-1].X]
			}
			matrix[pSet[0].Y][pSet[0].X] = tmp
		}
	}
}

// 计算某一个点的旋转集，n表示方阵维度
func calcRotateSet(pos Point2D, n int) []Point2D {
	ans := make([]Point2D, 4)
	ans[0] = pos
	for i := 1; i < 4; i++ {
		ans[i] = Point2D{
			X: n - ans[i-1].Y - 1,
			Y: ans[i-1].X,
		}
	}
	return ans
}
