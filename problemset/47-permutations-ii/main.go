package main

import (
	"fmt"
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []int{1, 1, 2}, Answer: [][]int{
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1},
		}},
		{Input: []int{1}, Answer: [][]int{{1}}},
		{Input: []int{}, Answer: [][]int{{}}},
		{Input: []int{1, 2, 3}, Answer: [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
		{Input: []int{2, 2, 3, 3}, Answer: [][]int{
			{2, 2, 3, 3},
			{3, 2, 2, 3},
			{3, 2, 3, 2},
			{3, 3, 2, 2},
			{2, 3, 2, 3},
			{2, 3, 3, 2},
		}},
	}

	test.Check(testCases,
		func(input interface{}) (result interface{}) {
			return permuteUnique(input.([]int))
		}, func(ans, res interface{}) bool {
			answer := ans.([][]int)
			result := res.([][]int)

			// 校验结果重复
			resultMap := make(map[string]bool)
			for _, r := range result {
				s := fmt.Sprintf("%v", r)
				if _, found := resultMap[s]; found {
					return false
				}
				resultMap[s] = true
			}

			// 校验结果正确性
			for _, a := range answer {
				aStr := fmt.Sprintf("%+v", a)
				matched := false
				for _, b := range result {
					bStr := fmt.Sprintf("%+v", b)
					if aStr == bStr {
						matched = true
						break
					}
				}
				if !matched {
					return false
				}
			}
			return true
		})
}

//----------------------------------------下面才是解题代码----------------------------------------

// 使用46题的代码算出有重复的结果集，然后再去重
func permuteUnique(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	ans := [][]int{}
	repeatableResult := permute(nums)

	// 结果去重
	resultMap := make(map[string][]int)
	for _, r := range repeatableResult {
		t := fmt.Sprintf("%v", r)
		if _, found := resultMap[t]; !found {
			resultMap[t] = r
		}
	}
	for _, v := range resultMap {
		ans = append(ans, v)
	}
	return ans
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		sub := append([]int{}, nums[:i]...)
		sub = append(sub, nums[i+1:]...)
		subAns := permute(sub)
		for _, s := range subAns {
			ans = append(ans, append([]int{nums[i]}, s...))
		}
	}
	return ans
}