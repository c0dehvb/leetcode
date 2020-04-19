package main

import "fmt"

func main() {
	ans := solveNQueens(4)
	for i, a := range ans {
		fmt.Printf("解法%d:\n", i+1)
		for _, row := range a {
			fmt.Println(row)
		}
		fmt.Println()
	}
}

//----------------------------------------下面才是解题代码----------------------------------------

var ans [][]string

func solveNQueens(n int) [][]string {
	ans = [][]string{}
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
		cp := make([]string, n)
		copy(cp, currentState)
		ans = append(ans, cp)
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
