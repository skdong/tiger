package hash

type Node struct {
	Value int
	Next  *Node
}

type MyHashSet struct {
	Cache []*Node
	Len   int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	len := 1000
	cache := make([]*Node, len)
	return MyHashSet{Cache: cache, Len: len}
}

func (this *MyHashSet) Add(key int) {
	node := this.Cache[key%this.Len]
	for node != nil {
		if node.Value == key {
			return
		}
		node = node.Next
	}
	node = this.Cache[key%this.Len]
	this.Cache[key%this.Len] = &Node{Value: key, Next: node}
}

func (this *MyHashSet) Remove(key int) {
	node := this.Cache[key%this.Len]
	last := &this.Cache[key%this.Len]
	for node != nil {
		if node.Value == key {
			*last = node.Next
		}
		last = &node.Next
		node = node.Next
	}

}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	node := this.Cache[key%this.Len]
	for node != nil {
		if node.Value == key {
			return true
		}
		node = node.Next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
