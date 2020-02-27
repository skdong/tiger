package queue

type StringNode struct {
	val  string
	next *StringNode
	last *StringNode
}

type StringQueue struct {
	head *StringNode
	tail *StringNode
}

func (q *StringQueue) empty() bool {
	return q.head == nil
}

func (q *StringQueue) add(s string) {
	node := &StringNode{val: s}
	if q.tail == nil {
		q.tail = node
		q.head = q.tail
	} else {
		q.tail.next = node
		node.last = q.tail
		q.tail = node
	}
}

func (q *StringQueue) pop() (s string) {
	if q.head == nil {
		return
	}
	s = q.head.val
	if q.head == q.tail {
		q.head, q.tail = nil, nil

	} else {
		q.head = q.head.next
		q.head.last = nil
	}
	return
}

func openLock(deadens []string, target string) int {
	step := 0
	tryed := map[string]bool{}
	for _, dead := range deadens {
		tryed[dead] = true
	}
	if _, ok := tryed[target]; ok {
		return -1
	}
	if _, ok := tryed["0000"]; ok {
		return -1
	}
	q := StringQueue{}
	q.add("0000")
	for !q.empty() {
		step++
		newq := StringQueue{}
		for !q.empty() {
			s := q.pop()
			if s == target {
				return step - 1
			}
			if _, ok := tryed[s]; ok {
				continue
			}
			tryed[s] = true
			for i := range s {
				b := []byte(s)
				if b[i] == '9' {
					b[i] = '8'
					newq.add(string(b))
					b[i] = '0'
					newq.add(string(b))
				} else if b[i] == '0' {
					b[i] = '9'
					newq.add(string(b))
					b[i] = '1'
					newq.add(string(b))
				} else {
					b[i] = b[i] + 1
					newq.add(string(b))
					b[i] = b[i] - 2
					newq.add(string(b))
				}
			}

		}
		q = newq
	}

	return -1
}
