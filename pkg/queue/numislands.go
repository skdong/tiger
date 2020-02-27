package queue

type Node struct {
	i    int
	j    int
	next *Node
	last *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) add(i, j int) {
	if q.head == nil {
		q.tail = &Node{i: i, j: j}
		q.head = q.tail
	} else {
		q.tail.next = &Node{i: i, j: j, last: q.tail}
		q.tail = q.tail.next
	}
}

func (q *Queue) pop() (i, j int) {
	if q.head != nil {
		i, j = q.head.i, q.head.j
		if q.head == q.tail {
			q.head, q.tail = nil, nil
		} else {
			q.head = q.head.next
			q.head.last = nil
		}
	} else {
		i, j = -1, -1
	}
	return
}

func (q *Queue) empty() bool {
	return q.head == nil
}

func numIslands(grid [][]byte) int {
	var num int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				num++
				q := Queue{}
				q.add(i, j)
				for !q.empty() {
					a, b := q.pop()
					if grid[a][b] == '0' {
						continue
					}
					grid[a][b] = '0'
					if a+1 < len(grid) && grid[a+1][b] == '1' {
						q.add(a+1, b)
					}
					if a-1 > -1 && grid[a-1][b] == '1' {
						q.add(a-1, b)
					}
					if b+1 < len(grid[0]) && grid[a][b+1] == '1' {
						q.add(a, b+1)
					}
					if b-1 > -1 && grid[a][b-1] == '1' {
						q.add(a, b-1)
					}
				}
			}
		}
	}
	return num
}
