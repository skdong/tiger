package queue

type NumNode struct {
	v    int
	next *NumNode
	last *NumNode
}

type NumQueue struct {
	head *NumNode
	tail *NumNode
}

func (q *NumQueue) empty() bool {
	return q.head == nil
}

func (q *NumQueue) add(n int) {
	node := &NumNode{v: n}
	if q.tail == nil {
		q.tail, q.head = node, node
	} else {
		node.next = q.head
		q.head.last = node
		q.head = node
	}
}
func (q *NumQueue) pop() int {
	if q.tail == nil {
		return -1
	}
	v := q.tail.v
	if q.tail == q.head {
		q.tail, q.head = nil, nil
	} else {
		q.tail = q.tail.last
		q.tail.next = nil
	}
	return v

}

func numSquares(n int) int {
	cache := map[int]bool{}
	q := NumQueue{}
	q.add(n)
	var num int
	for !q.empty() {
		nq := NumQueue{}
		num++
		for !q.empty() {
			m := q.pop()
			if _, ok := cache[m]; ok {
				continue
			}
			cache[m] = true
			for i := 0; i*i <= m; i++ {
				if m-i*i == 0 {
					return num
				}
				nq.add(m - i*i)
			}
		}
		q = nq
	}
	return num
}

func numSquares1(n int) int {
	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		a[i] = i
		for j := 1; i-j*j >= 0; j++ {
			if a[i] > a[i-j*j]+1 {
				a[i] = a[i-j*j] + 1
			}
		}
	}
	return a[n]
}
