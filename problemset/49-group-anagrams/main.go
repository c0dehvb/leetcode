package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
	"sort"
)

func main() {
	testCases := []test.TestCase{
		{Input: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, Answer: [][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		}},
		{Input: []string{"eat"}, Answer: [][]string{
			{"eat"},
		}},
		{Input: []string{}, Answer: [][]string{
		}},
	}

	test.SimpleCheck(testCases,
		func(input interface{}) (result interface{}) {
			res := groupAnagrams(input.([]string))
			for i, _ := range res {
				sort.Strings(res[i])
			}
			return res
		})
}

//----------------------------------------下面才是解题代码----------------------------------------

// 同一个异位词集中的所有字符串将其自身字符按字典序排序后得出来的新串是一样的
// 把新串作为key，用map做归类即可
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, s := range strs {
		key := sortChar(s)
		value, found := m[key]
		if !found {
			value = []string{}
		}
		value = append(value, s)
		m[key] = value
	}

	var ans [][]string
	for _, valueSet := range m {
		ans = append(ans, valueSet)
	}
	return ans
}

// 将一个字符串中的字符按字典序排序
func sortChar(str string) string {
	nums := make([]int, len(str))
	for i, c := range str {
		nums[i] = int(c)
	}
	sort.Ints(nums)
	ans := ""
	for _, i := range nums {
		ans += string(i)
	}
	return ans
}
