package container

import (
	"testing"
)

func TestQueue(t *testing.T) {
	items := []Value{1, 2, 3}
	target := []int{}
	queue := CreateQueue()
	queue.Add(items...)
	for !queue.Empty() {
		val := queue.Pop()
		target = append(target, val.(int))
	}
	for i := range target {
		if items[i].(int) != target[i] {
			t.Fatalf("")
		}
	}
}
