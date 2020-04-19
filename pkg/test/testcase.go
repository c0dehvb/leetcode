package test

import "fmt"

type TestCase struct {
	Input  interface{}
	Answer interface{}
}

func SimpleCheck(testCases []TestCase, solve func(input interface{}) (result interface{})) {
	Check(testCases, solve, func(answer, result interface{}) bool {
		resStr := fmt.Sprintf("%+v", result)
		ansStr := fmt.Sprintf("%+v", answer)
		return ansStr == resStr
	})
}

func Check(testCases []TestCase,
	solve func(input interface{}) (result interface{}),
	matchFunc func(answer, result interface{}) bool) {

	for i, t := range testCases {
		result := solve(t.Input)
		if !matchFunc(t.Answer, result) {
			panic(fmt.Errorf("第%d个测试错误: 答案=%v, 计算结果=%v", i+1, t.Answer, result))
		}
	}

	fmt.Printf("%d个测试用例全部通过!", len(testCases))
}
