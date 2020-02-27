package container

type Value interface{}

type Node struct {
	Val  Value
	Next *Node
	Last *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func CreateQueue() Queue {
	return Queue{}
}

func (q *Queue) Add(vals ...Value) {
	for _, val := range vals {
		q.add(val)
	}
}

func (q *Queue) add(val Value) {
	node := &Node{Val: val, Next: q.head}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.head.Last = node
		q.head = node
	}
}

func (q *Queue) Pop() (val Value) {
	if q.tail != nil {
		val = q.tail.Val
		if q.head == q.tail {
			q.head = nil
			q.tail = nil
		} else {
			q.tail = q.tail.Last
			q.tail.Next = nil
		}
	}
	return
}

func (q *Queue) Empty() bool {
	return q.head == nil
}
