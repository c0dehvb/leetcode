package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}, Answer: [][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}},
		{Input: [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}, Answer: [][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		matrix := input.([][]int)
		setZeroes(matrix)
		return matrix
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func setZeroes(matrix [][]int) {
	num := matrix[0][0] + 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if num < matrix[i][j]+1 {
				num = matrix[i][j] + 1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				setNums(matrix, i, j, num)
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == num {
				matrix[i][j] = 0
			}
		}
	}
}

func setNums(matrix [][]int, x, y, num int) {
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[x][i] != 0 {
			matrix[x][i] = num
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][y] != 0 {
			matrix[i][y] = num
		}
	}
}
