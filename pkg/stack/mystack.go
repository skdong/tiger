package stack

type MinStack struct {
	s []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
}

func (this *MinStack) Pop() {
	if len(this.s) > 0 {
		this.s = this.s[1:]
	}
}

func (this *MinStack) Top() int {
	if len(this.s) > 0 {
		return this.s[0]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.s) < 1 {
		return -1
	}
	min := this.s[0]
	for i := 1; i < len(this.s); i++ {
		if min > this.s[i] {
			min = this.s[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
