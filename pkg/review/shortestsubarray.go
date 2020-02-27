package review

type Node struct {
	Value int
	Next  *Node
	Last  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Empty() bool {
	return q.Head == nil
}

func (q *Queue) Push(value int) {
	if q.Empty() {
		q.Tail = &Node{Value: value}
		q.Head = q.Tail
	} else {
		q.Tail.Next = &Node{Value: value, Last: q.Tail}
		q.Tail = q.Tail.Next
	}
}

func (q *Queue) Front() (value int) {
	if !q.Empty() {
		value = q.Head.Value
	}
	return
}

func (q *Queue) Back() (value int) {
	if !q.Empty() {
		value = q.Tail.Value
	}
	return
}

func (q *Queue) PopFront() (value int) {
	if !q.Empty() {
		value = q.Head.Value
		if q.Head == q.Tail {
			q.Head = nil
			q.Tail = nil
		} else {
			q.Head = q.Head.Next
			q.Head.Last = nil
		}
	}
	return
}

func (q *Queue) PopBack() (value int) {
	if !q.Empty() {
		value = q.Tail.Value
		if q.Head == q.Tail {
			q.Head = nil
			q.Tail = nil
		} else {
			q.Tail = q.Tail.Last
			q.Tail.Next = nil
		}
	}
	return
}

func shortestSubarray(A []int, K int) int {
	shortest := -1
	minQueue := Queue{}
	sumArray := make([]int, len(A)+1)
	for i := 0; i < len(A); i++ {
		sumArray[i+1] = sumArray[i] + A[i]
	}
	for i := 0; i < len(sumArray); i++ {
		for !minQueue.Empty() && sumArray[i] < sumArray[minQueue.Back()] {
			minQueue.PopBack()
		}
		for !minQueue.Empty() && sumArray[i]-sumArray[minQueue.Front()] >= K {
			if shortest == -1 || i-minQueue.Front() < shortest {
				shortest = i - minQueue.Front()
			}
			minQueue.PopFront()
		}
		minQueue.Push(i)
	}
	return shortest
}
