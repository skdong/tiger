package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func test() {
	rand.Seed(time.Now().UnixNano())
	n := 10
	items := make([]int, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, rand.Intn(n))

	}
	fmt.Println(items)
}

func producter(wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 10; i++ {
		item := rand.Intn(20)

		ch <- item
	}
	ch <- -1
	defer wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for {
		item := <-ch
		if item == -1 {
			fmt.Println("over")
			return
		}
		fmt.Println(item)
	}
}

func ProductConsumer() {

	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}

	ch := make(chan int, 100)

	wg.Add(1)
	go producter(&wg, ch)

	wg.Add(1)
	go consumer(&wg, ch)

	wg.Wait()
}

func testCh() {
	ch := make(chan int)
	stopCh := make(chan int)
	select {
	case <-stopCh:
		fmt.Println("stoped")
	case item := <-ch:
		fmt.Println(item)
	}
}

type Caches struct {
	Lock  sync.Mutex
	store map[string]string
	back  map[string]string
}

func (c *Caches) Get(key string) string {
	value, err := c.store[key]
	if !err {
		fmt.Printf("has no key [%v]\n", key)
	}
	return value
}

func (c *Caches) Set(key, value string) {
	c.store[key] = value
}

func (c *Caches) Delete(key string) {
	delete(c.store, key)
}

func testCahces() {
	cache := Caches{store: make(map[string]string)}
	cache.Get("test")
	cache.Set("test", "aa")
	fmt.Println(cache.Get("test"))
	cache.Delete("test")
	fmt.Println(cache.Get("test"))
	fmt.Println(cache.Get("test"))
}

var cache [][]int

func isMatchBytes(i, j int, s []byte, p []byte) bool {
	if cache[i][j] != 0 {
		if cache[i][j] == 1 {
			return true
		} else {
			return false
		}
	}
	if len(p) == 0 {
		if len(s) == 0 {
			cache[i][j] = 1
			return true
		} else {
			cache[i][j] = 2
			return false
		}
	}
	if len(p) > 1 && p[1] == '*' {
		if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
			if isMatchBytes(i+1, j, s[1:], p) {
				cache[i][j] = 1
				return true
			} else if isMatchBytes(i, j+2, s, p[2:]) {
				cache[i][j] = 1
				return true
			} else {
				cache[i][j] = 2
				return false
			}
		} else {
			if isMatchBytes(i, j+2, s, p[2:]) {
				cache[i][j] = 1
				return true
			} else {
				cache[i][j] = 2
				return false
			}
		}
	} else {
		if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
			if isMatchBytes(i+1, j+1, s[1:], p[1:]) {
				cache[i][j] = 1
				return true
			} else {
				cache[i][j] = 2
				return false
			}
		} else {
			cache[i][j] = 2
			return false
		}
	}
}

func isMatch(s string, p string) bool {
	cache = make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		cache[i] = make([]int, len(p))
	}

	return isMatchBytes(0, 0, []byte(s), []byte(p))
}

func testIsMatch() {
	fmt.Println(isMatch("aa", "a*"))
}

func testBytes() {
	s := "test"
	s_bytes := []byte(s)
	fmt.Println(s_bytes)
}

func testList() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	b := a
	b[0] = 2
	fmt.Println(a)
	fmt.Println(a[1:3])
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func LoadTree() {
	tree := []int{5, 1, 4, -1, -1, 3, 6}
	tmp := make([]Node, len(tree))
	for i := 0; i < len(tree); i++ {
		if tree[i] == -1 {
			continue
		}
		node := Node{Val: tree[i]}
		tmp[i] = node
		if i > 0 {
			if i-1%2 == 0 {
				tmp[(i-1)>>1].Left = &node
			} else {
				tmp[(i-1)>>1].Right = &node
			}
		}
	}
}

func testAlise() {
	a := 10
	b := &a
	c := &b
	*b = 2
	fmt.Println(a)
	*c = nil
	fmt.Println(b)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func testTree() {
	root := TreeNode{Val: 1}
	node := &root
	fmt.Println((*node).Left)
}

func reverseString(s []byte) {

}

func main() {

	reverseString([]byte("hello"))
}
