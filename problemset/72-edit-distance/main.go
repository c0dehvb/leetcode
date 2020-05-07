package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []string{"sea", "eat"}, Answer: 2},
		{Input: []string{"horse", "ros"}, Answer: 3},
		{Input: []string{"intention", "execution"}, Answer: 5},
		{Input: []string{"", "aaaaa"}, Answer: 5},
		{Input: []string{"aaaaa", ""}, Answer: 5},
		{Input: []string{"adada", "aaa"}, Answer: 2},
		{Input: []string{"adada", "ddd"}, Answer: 3},
		{Input: []string{"dinitrophenylhydrazine", "acetylphenylhydrazine"}, Answer: 6},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return minDistance(input.([]string)[0], input.([]string)[1])
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word1)+2)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = len(word1) + len(word2) + 10
		}
	}
	for j := 0; j <= len(word1); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word2); i++ {
		for j := 1; j <= len(word1)+1; j++ {
			if j < len(word1)+1 {
				// 有相同字母则直接使用
				if word2[i-1] == word1[j-1] {
					dp[i][j] = MinInt(dp[i-1][j-1], dp[i][j])
				}
				// 替换掉第j个字母为问号
				dp[i][j] = MinInt(dp[i-1][j-1]+1, dp[i][j])
				// 在第j个字母前插入问号
				dp[i][j-1] = MinInt(dp[i-1][j-1]+1, dp[i][j-1])

				// 后面多出来的当作删除操作
				for k := j + 1; k <= len(word1); k++ {
					dp[i][k] = MinInt(dp[i][j]+k-j, dp[i][k])
				}
			} else { // 在末尾插入问号
				dp[i][j] = MinInt(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j])
			}
		}
	}
	return MinInt(dp[len(dp)-1][len(dp[0])-1], dp[len(dp)-1][len(dp[0])-2])
}

func MinInt(nums ...int) int {
	ans := nums[0]
	for _, n := range nums {
		if n < ans {
			ans = n
		}
	}
	return ans
}
