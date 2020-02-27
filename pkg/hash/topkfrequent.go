package hash

import (
	"strconv"
	"strings"
)

type NumCounter struct {
	Num     int
	Counter int
}

type Bucket struct {
	Nums []*NumCounter
	k    int
	i    int
}

func CreateBucket(top int) *Bucket {
	return &Bucket{Nums: make([]*NumCounter, top+1), k: top}
}

func (b *Bucket) String() string {
	s := make([]string, b.i)
	for i := range s {
		s[i] = strconv.Itoa(b.Nums[i].Num)
	}
	return strings.Join(s, ",")
}

func (b *Bucket) rootAdjust() {
	b.Nums[0] = b.Nums[b.i]
	curr := 0
	next := curr*2 + 1
	for next < b.i-1 {
		if next+1 <= b.i-1 && b.Nums[next].Counter > b.Nums[next+1].Counter {
			next++
		}
		if b.Nums[next].Counter >= b.Nums[curr].Counter {
			break
		}
		b.Nums[next], b.Nums[curr] = b.Nums[curr], b.Nums[next]
		curr = next
		next = curr*2 + 1
	}
}

func (b *Bucket) adjust() {
	curr := b.i
	for curr > 0 {
		uper := (curr - 1) / 2
		if b.Nums[uper].Counter > b.Nums[curr].Counter {
			b.Nums[uper], b.Nums[curr] = b.Nums[curr], b.Nums[uper]
			curr = uper
		} else {
			break
		}

	}
}

func (b *Bucket) Add(n *NumCounter) {
	if b.i == b.k && b.Nums[0].Counter >= n.Counter {
		return
	}
	b.Nums[b.i] = n
	b.adjust()
	if b.i < b.k {
		b.i++
	} else {
		b.rootAdjust()
	}
}

func (b *Bucket) Get() []int {
	topK := make([]int, b.k)
	for i := range topK {
		topK[i] = b.Nums[i].Num
	}
	return topK
}

func topKFrequent(nums []int, k int) []int {
	counter := map[int]int{}
	for _, n := range nums {
		if _, ok := counter[n]; !ok {
			counter[n] = 1
		} else {
			counter[n] += 1
		}
	}
	bucket := CreateBucket(k)
	for num, counter := range counter {
		bucket.Add(&NumCounter{Num: num, Counter: counter})
	}
	return bucket.Get()
}
