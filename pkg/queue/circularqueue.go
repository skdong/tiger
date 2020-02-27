package queue

type MyCircularQueue struct {
	head  int
	tail  int
	array []int
	k     int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{array: make([]int, k),
		k:    k,
		head: -1,
		tail: -1}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.head != -1 && this.head == this.tail {
		return false
	}
	if this.tail == -1 {
		this.head = 0
		this.tail = 0
	}
	this.array[this.tail] = value
	this.tail = (this.tail + 1) % this.k
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.tail == -1 {
		return false
	}
	this.head = (this.head + 1) % this.k
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
	}
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.head != -1 {
		return this.array[this.head]
	}
	return -1
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.head != -1 {
		return this.array[(this.tail-1+this.k)%this.k]
	}
	return -1
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == -1
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.head != -1 && this.head == this.tail
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
