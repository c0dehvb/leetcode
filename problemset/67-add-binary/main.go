package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []string{"11", "1"}, Answer: "100"},
		{Input: []string{"1010", "1011"}, Answer: "10101"},
		{Input: []string{"0", "1011"}, Answer: "1011"},
		{Input: []string{"1010", "0"}, Answer: "1010"},
		{Input: []string{"0", "0"}, Answer: "0"},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return addBinary(input.([]string)[0], input.([]string)[1])
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func addBinary(a string, b string) string {
	var aNum, bNum byte
	ans := []byte{}
	curIdx := 0
	plusOne := 0

	for curIdx < len(a) || curIdx < len(b) {
		if curIdx < len(a) {
			aNum = a[len(a)-curIdx-1] - '0'
		} else {
			aNum = 0
		}
		if curIdx < len(b) {
			bNum = b[len(b)-curIdx-1] - '0'
		} else {
			bNum = 0
		}
		ans = append([]byte{(aNum + bNum + byte(plusOne)) % 2}, ans...)
		if aNum+bNum+byte(plusOne) > 1 {
			plusOne = 1
		} else {
			plusOne = 0
		}
		curIdx++
	}
	if plusOne == 1 {
		ans = append([]byte{1}, ans...)
	}
	ansStr := ""
	for _, b := range ans {
		ansStr += string(b + '0')
	}
	return ansStr
}
