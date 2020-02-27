package set

type RandomizedSet struct {
	cache  [][]int
	length int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{cache: make([][]int, 1000), length: 1000}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if this.cache[val%this.length] == nil {
		this.cache[val%this.length] = []int{val}
	}else{
		for _, v := this.cache[val%this.length]{
			if v == val{
				return false
			}
		}
	}

}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {

}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {

}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
