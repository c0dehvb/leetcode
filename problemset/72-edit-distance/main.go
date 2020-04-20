package main

import (
	"fmt"
	"github.com/c0dehvb/leetcode/pkg/test"
)

func main() {
	testCases := []test.TestCase{
		{Input: []string{"horse", "ros"}, Answer: 3},
		{Input: []string{"intention", "execution"}, Answer: 5},
	}

	test.SimpleCheck(testCases, func(input interface{}) (result interface{}) {
		return minDistance(input.([]string)[0], input.([]string)[1])
	})
}

//----------------------------------------下面才是解题代码----------------------------------------

type Item struct {
	word  string
	score int
}

type Queue struct {
	elems []Item
	l, r  int
}

var (
	ErrQueueEmpty = fmt.Errorf("queue is empty")
)

func NewQueue(cap int) *Queue {
	return &Queue{
		elems: make([]Item, cap),
		l:     0,
		r:     0,
	}
}

func (q *Queue) Push(d Item) {
	if q.r < len(q.elems) {
		q.elems[q.r] = d
		q.r++
	} else {
		q.elems = append(q.elems, make([]Item, len(q.elems))...)
		q.Push(d)
	}
}

func (q *Queue) Pop() (Item, error) {
	p := Item{}
	if q.Size() == 0 {
		return p, ErrQueueEmpty
	} else {
		p = q.elems[q.l]
		q.l++
		return p, nil
	}
}

func (q *Queue) Size() int {
	return q.r - q.l
}

var ans int

// 推论1：两个单词变换次数不大于max(len(word1), len(word2))
func minDistance(word1 string, word2 string) int {
	dist := make(map[string]int)
	dist[word1] = 0
	initScore := score(word1, word2)
	if initScore == 0 {
		return 0
	}
	ans = initScore
	queue := NewQueue(initScore + 1)
	queue.Push(Item{word1, initScore})
	for queue.Size() > 0 {
		head, _ := queue.Pop()
		initScore = head.score
		word := head.word

		// replace
		for i := 0; i < len(word); i++ {
			newWord := replace(word, i)
			dealNewWord(newWord, word, word2, dist, initScore, queue)
		}

		// insert
		for i := 0; i <= len(word); i++ {
			newWord := insert(word, i)
			dealNewWord(newWord, word, word2, dist, initScore, queue)
		}

		// delete
		for i := 0; i < len(word); i++ {
			newWord := remove(word, i)
			dealNewWord(newWord, word, word2, dist, initScore, queue)
		}
	}
	return ans
}

func dealNewWord(newWord, srcWord, word2 string, dist map[string]int, initScore int, queue *Queue) {
	if _, found := dist[newWord]; !found || dist[newWord] > dist[srcWord]+1 {
		dist[newWord] = dist[srcWord] + 1
		newScore := score(newWord, word2)
		if newScore < initScore && newScore > 0 {
			queue.Push(Item{newWord, newScore})
		} else if newScore == 0 && dist[newWord] < ans {
			ans = dist[newWord]
		}
	}
}

func score(word1 string, word2 string) int {
	ans := 0
	s1Idx := 0
	s2Idx := 0
	for s1Idx < len(word1) && s2Idx < len(word2) {
		if word1[s1Idx] == word2[s2Idx] || word1[s1Idx] == '?' {
			s2Idx++
		} else {
			ans++
		}
		s1Idx++
	}
	ans += len(word2) - s2Idx + len(word1) - s1Idx
	return ans
}

func insert(word string, idx int) string {
	if idx < len(word) {
		return word[:idx] + "?" + word[idx:]
	} else {
		return word + "?"
	}
}

func remove(word string, idx int) string {
	if idx < len(word)-1 {
		return word[:idx] + word[idx+1:]
	} else {
		return word[:idx]
	}
}

func replace(word string, i int) string {
	b := []byte(word)
	b[i] = '?'
	return string(b)
}
