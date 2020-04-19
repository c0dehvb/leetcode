package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []interface{}{"acb", 4, "ab", 2}, Answer: 2},
		{Input: []interface{}{"bca", 4, "ab", 2}, Answer: 1},
		{Input: []interface{}{"bca", 4, "acb", 2}, Answer: 0},
		{Input: []interface{}{"bca", 5, "acb", 2}, Answer: 1},
		{Input: []interface{}{"aaa", 3, "aa", 1}, Answer: 4},
		{Input: []interface{}{"aaa", 1, "aa", 2}, Answer: 0},
		{Input: []interface{}{"aaa", 0, "aa", 1}, Answer: 0},
		{Input: []interface{}{"lovelive", 100000, "lovelive", 100000}, Answer: 1},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return getMaxRepetitions(input.([]interface{})[0].(string), input.([]interface{})[1].(int), input.([]interface{})[2].(string), input.([]interface{})[3].(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

type Record struct {
	n1                    int // 要多少个s1串才能组成一个s2串
	remainderStringLength int // 拼完s2串后剩余最后一个s1串的长度，用于接着拼下一个s2串
}

// 数据范围: s1,s2串长度最大100，0<=n1,n2<=10^6
// 推论1: s1串能组成s2串的充要条件是s2串上每一个字符都必须存在于s1串上。
// 推论2: 如果s1有能力组成s2串，那么组成一个s2串最多只需要len(s2)个s1串，即不大于100个。
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	s1CharMap := make(map[uint8]bool)
	s2CharMap := make(map[uint8]bool)
	// 验证推论1
	for i, _ := range s1 {
		s1CharMap[s1[i]] = true
	}
	for i, _ := range s2 {
		s2CharMap[s2[i]] = true
		if _, found := s1CharMap[s2[i]]; !found {
			return 0
		}
	}

	// 初始化records数组
	// 数组下标表示前一次拼完s2串后剩余s1串的长度
	// records[i]指以s1[i+1:]开始往后继续拼多少个完整s1串才能形成一个s2串
	records := make([]Record, len(s1)+1)
	records[len(s1)] = buildRecord(s1, s2, len(s1))
	records[0] = records[len(s1)]
	records[0].n1++
	for i := 1; i < len(s1); i++ {
		if s1[len(s1)-i] != s2[0] {
			records[i] = records[i-1]
		} else {
			records[i] = buildRecord(s1, s2, i)
		}
	}

	if records[len(s1)].n1 > n1 { // 给出的n1不足以拼成1个s2
		return 0
	} else {
		ans := 0
		rsl := len(s1)
		for n1 >= records[rsl].n1 {
			ans++
			n1 -= records[rsl].n1
			// 如果有字符剩余可以留给下一次使用
			if records[rsl].remainderStringLength > 0 {
				n1++
				rsl = records[rsl].remainderStringLength
			} else {
				rsl = len(s1)
			}
		}
		return ans / n2
	}
}

func buildRecord(s1 string, s2 string, remainderStringLength int) Record {
	ans := 1
	s1Idx := len(s1) - remainderStringLength
	s2Idx := 0

	for s2Idx < len(s2) {
		if s1[s1Idx] == s2[s2Idx] {
			s2Idx++
		}
		s1Idx++
		if s2Idx < len(s2) && s1Idx == len(s1) {
			s1Idx = 0
			ans++
		}
	}
	return Record{ans, len(s1) - s1Idx}
}
