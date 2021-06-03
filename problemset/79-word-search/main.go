package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []interface{}{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"}, Answer: true},
		{Input: []interface{}{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"}, Answer: true},
		{Input: []interface{}{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"}, Answer: false},
		{Input: []interface{}{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "A"}, Answer: true},
		{Input: []interface{}{[][]byte{{'A'}}, "A"}, Answer: true},
		{Input: []interface{}{[][]byte{{'B'}}, "A"}, Answer: false},
		{Input: []interface{}{[][]byte{{'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}}, "FFFFFFE"}, Answer: true},
		{Input: []interface{}{[][]byte{{'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}, {'A', 'B', 'C', 'D', 'E', 'F'}}, "FFFFFFD"}, Answer: false},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return exist(input.([]interface{})[0].([][]byte), input.([]interface{})[1].(string))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

var visited [][]bool

// 方向，上下左右
var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

// m = board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			// 找到开头字母配对的格子开始搜索
			if board[x][y] == word[0] {
				visited[x][y] = true
				if search(x, y, board, word[1:]) {
					return true
				}
				visited[x][y] = false
			}
		}
	}
	return false
}

func search(x, y int, board [][]byte, word string) bool {
	if word == "" { // 到底了，找到答案
		return true
	}

	// 每个方向都尝试一下
	for i := 0; i < len(dir); i++ {
		tx := x + dir[i][0]
		ty := y + dir[i][1]

		// 判断是否跑出边界
		if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
			continue
		}
		// 不走重复走过的路
		if visited[tx][ty] {
			continue
		}

		// 确定字符符合单词顺序才可以往下搜索
		if board[tx][ty] == word[0] {
			visited[tx][ty] = true
			if search(tx, ty, board, word[1:]) {
				return true
			}
			visited[tx][ty] = false
		}
	}
	return false // 搜索完所有方向都没找到
}
