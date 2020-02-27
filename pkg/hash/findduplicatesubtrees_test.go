package hash

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func loadTree() *TreeNode {
	var root *TreeNode
	fileName := "C:\\Users\\kaidong\\Documents\\tmp\\tree"
	f, err := os.Open(fileName)
	if err != nil {
		return root
	}
	defer f.Close()
	info, _ := ioutil.ReadAll(f)
	items := strings.Split(string(info), ",")
	return transItemsToTreeNode(items)
}

func transItemsToTreeNode(items []string) *TreeNode {
	var root *TreeNode
	val, err := strconv.Atoi(items[0])
	root = &TreeNode{Val: val}
	queue := CreateQueue()
	queue.add(root)
	for i := 1; i < len(items); i += 2 {
		father := queue.pop()
		val, err = strconv.Atoi(items[i])
		if err == nil {
			node := &TreeNode{Val: val}
			father.Left = node
			queue.add(node)
		}
		if i+1 < len(items) {
			val, err = strconv.Atoi(items[i+1])
			if err == nil {
				node := &TreeNode{Val: val}
				father.Right = node
				queue.add(node)
			}

		}
	}
	return root
}

func TestFindDuplicateSubtrees(t *testing.T) {
	trees := findDuplicateSubtrees(loadTree())
	t.Fatal(trees)
}
