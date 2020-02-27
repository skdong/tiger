package hash

type MapNode struct {
	Key   int
	Value int
	Next  *MapNode
}
type MyHashMap struct {
	Len   int
	Cache []*MapNode
}

/** Initialize your data structure here. */
func ConstructorMap() MyHashMap {
	len := 10000
	return MyHashMap{Len: len, Cache: make([]*MapNode, len)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	node := this.Cache[key%this.Len]
	for node != nil {
		if node.Key == key {
			node.Value = value
			return
		}
		node = node.Next
	}
	node = this.Cache[key%this.Len]
	this.Cache[key%this.Len] = &MapNode{Key: key,
		Value: value,
		Next:  node}

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	node := this.Cache[key%this.Len]
	for node != nil {
		if node.Key == key {
			return node.Value
		}
		node = node.Next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	last := &this.Cache[key%this.Len]
	node := *last
	for node != nil {
		if node.Key == key {
			*last = node.Next
		}
		last = &node.Next
		node = *last
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
