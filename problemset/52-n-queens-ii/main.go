package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: 4, Answer: 2},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return totalNQueens(input.(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

var ans int

// 把第52题的代码拿过来小改就好了
func totalNQueens(n int) int {
	ans = 0
	row := ""
	for i := 0; i < n; i++ {
		row += "."
	}
	currentState := make([]string, n)
	for i := 0; i < n; i++ {
		currentState[i] = row
	}
	solveNQueensInternal(0, n, currentState)
	return ans
}

func solveNQueensInternal(deep, n int, currentState []string) {
	if deep == n {
		ans++
		return
	}
Loop:
	for i := 0; i < n; i++ {
		// check
		for j := 0; j < deep; j++ {
			if currentState[j][i] == 'Q' || // 竖线
				i-(deep-j) >= 0 && currentState[j][i-(deep-j)] == 'Q' || // 左斜线
				i+(deep-j) < n && currentState[j][i+(deep-j)] == 'Q' { // 右斜线
				continue Loop
			}
		}
		currentState[deep] = currentState[deep][:i] + "Q" + currentState[deep][i+1:]
		solveNQueensInternal(deep+1, n, currentState)
		currentState[deep] = currentState[deep][:i] + "." + currentState[deep][i+1:]
	}
}
