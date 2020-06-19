package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []interface{}{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			3,
		}, Answer: true},
		{Input: []interface{}{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			13,
		}, Answer: false},
		{Input: []interface{}{
			[][]int{
				{1},
			},
			1,
		}, Answer: true},
		{Input: []interface{}{
			[][]int{
				{1},
			},
			2,
		}, Answer: false},
		{Input: []interface{}{
			[][]int{
				{1, 3},
			},
			3,
		}, Answer: true},
		{Input: []interface{}{
			[][]int{
				{1, 3},
			},
			1,
		}, Answer: true},
		{Input: []interface{}{
			[][]int{
				{1, 3},
			},
			2,
		}, Answer: false},
		{Input: []interface{}{
			[][]int{
				{1},
				{3},
			},
			1,
		}, Answer: true},
		{Input: []interface{}{
			[][]int{
				{1},
				{3},
			},
			3,
		}, Answer: true},
		{Input: []interface{}{
			[][]int{
				{1},
				{3},
			},
			2,
		}, Answer: false},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return searchMatrix(input.([]interface{})[0].([][]int), input.([]interface{})[1].(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// 竖着二分
	var l, r, tmp, y int
	l = 0
	r = len(matrix) - 1
	for l <= r {
		tmp = matrix[(l+r)/2][0]
		if tmp == target {
			return true
		} else if tmp > target {
			r = (l+r)/2 - 1
		} else {
			l = (l+r)/2 + 1
		}
	}

	if r < 0 {
		return false
	}
	y = l - 1

	// 横着二分
	l = 0
	r = len(matrix[y]) - 1
	for l <= r {
		tmp = matrix[y][(l+r)/2]
		if tmp == target {
			return true
		} else if tmp > target {
			r = (l+r)/2 - 1
		} else {
			l = (l+r)/2 + 1
		}
	}
	return false
}
