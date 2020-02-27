package hash

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Node *TreeNode
	Next *ListNode
	Last *ListNode
}

type Queue struct {
	head *ListNode
	tail *ListNode
}

func CreateQueue() Queue {
	return Queue{}
}

func (q *Queue) add(nodes ...*TreeNode) {
	for _, node := range nodes {
		if q.head == nil {
			q.head = &ListNode{Node: node}
			q.tail = q.head
		} else {
			q.tail.Next = &ListNode{Node: node, Last: q.tail}
			q.tail = q.tail.Next
		}
	}
}

func (q *Queue) empty() bool {
	return q.head == nil
}

func (q *Queue) pop() *TreeNode {
	var node *TreeNode
	if q.head == nil {
		return node
	} else {
		node = q.head.Node
		if q.head == q.tail {
			q.head = nil
			q.tail = nil
		} else {
			q.head = q.head.Next
			q.head.Last = nil
		}
	}
	return node
}

type Manager struct {
	Trees []*TreeNode
	Cache map[string]bool
}

func (m *Manager) treeToString(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	tree := fmt.Sprintf("%s(%s,%s)",
		strconv.Itoa(root.Val),
		m.treeToString(root.Left),
		m.treeToString(root.Right))

	if r, ok := m.Cache[tree]; ok {
		if r {
			m.Trees = append(m.Trees, root)
			m.Cache[tree] = false
		}
	} else {
		m.Cache[tree] = true
	}
	return tree
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	manager := Manager{Cache: map[string]bool{}}
	manager.treeToString(root)

	return manager.Trees
}

func treeTolist(root *TreeNode) string {
	var tree string
	queue := CreateQueue()
	queue.add(root)
	for !queue.empty() {
		node := queue.pop()
		if node != nil {
			queue.add(node.Left, node.Right)
			tree = tree + "," + string(node.Val)
		} else {
			tree = tree + ",nil"
		}
	}
	return tree
}
