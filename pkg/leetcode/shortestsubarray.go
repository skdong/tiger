package leetcode

import (
	"container/list"
)

func shortestSubArray(A []int, K int) int {
	var p []int = []int{0}
	var ans int = len(A) + 1

	que := list.New()
	for i := range A {
		p = append(p, p[len(p)-1]+A[i])
	}
	for i := range p {
		for que.Len() > 0 && p[i] < p[que.Back().Value.(int)] {
			que.Remove(que.Back())
		}

		for que.Len() > 0 && p[i]-p[que.Front().Value.(int)] >= K {
			if ans > i-que.Front().Value.(int) {
				ans = i - que.Front().Value.(int)
			}
			que.Remove(que.Front())
		}
		que.PushBack(i)
	}
	if ans == len(A)+1 {
		return -1
	} else {
		return ans
	}
}
