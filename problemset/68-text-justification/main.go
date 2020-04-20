package main

import (
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []interface{}{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16}, Answer: []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		}},
		{Input: []interface{}{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16}, Answer: []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}},
		{Input: []interface{}{[]string{"Science", "is", "what", "we", "understand", "well", "enough",
			"to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20}, Answer: []string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		}},
		{Input: []interface{}{[]string{"hello", "golang,", "nice", "to", "meet", "you."}, 7}, Answer: []string{
			"hello  ",
			"golang,",
			"nice to",
			"meet   ",
			"you.   ",
		}},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return fullJustify(input.([]interface{})[0].([]string), input.([]interface{})[1].(int))
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

func fullJustify(words []string, maxWidth int) []string {
	if len(words) == 0 {
		return words
	}
	var ans []string
	var line string
	idx := 0
	for idx < len(words) {
		line, idx = fetchLine(words, maxWidth, idx)
		ans = append(ans, line)
	}
	return ans
}

func fetchLine(words []string, maxWidth, idx int) (string, int) {
	fetchWords := []string{}
	length := 0
	wordLength := 0
	for ; idx < len(words) && length < maxWidth; idx++ {
		if length == 0 {
			fetchWords = append(fetchWords, words[idx])
			length += len(words[idx])
			wordLength += len(words[idx])
		} else if length+len(words[idx])+1 <= maxWidth {
			fetchWords = append(fetchWords, words[idx])
			length += len(words[idx]) + 1
			wordLength += len(words[idx])
		} else {
			break
		}
	}

	var ans string
	if idx == len(words) { // 最后一行应为左对齐，且单词之间不插入额外的空格。
		for i, word := range fetchWords {
			if i == 0 {
				ans = word
			} else {
				ans += " " + word
			}
		}
		ans += nSpace(maxWidth - wordLength - (len(fetchWords) - 1))
	} else if len(fetchWords) == 1 { // 只包含一个单词，左对齐
		ans = fetchWords[0]
		ans += nSpace(maxWidth - wordLength)
	} else { // 普通行，应保证空格尽可能均匀
		wordCount := len(fetchWords)
		spaceCount := maxWidth - wordLength
		overflowCount := spaceCount % (wordCount - 1) // 如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数
		for i, word := range fetchWords {
			if i == 0 {
				ans = word
			} else if i <= overflowCount {
				ans += nSpace(spaceCount/(wordCount-1)+1) + word
			} else {
				ans += nSpace(spaceCount/(wordCount-1)) + word
			}
		}
	}
	return ans, idx
}

func nSpace(n int) string {
	ans := ""
	for i := 0; i < n; i++ {
		ans += " "
	}
	return ans
}
