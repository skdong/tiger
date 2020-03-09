package leetcode

import "fmt"

func numberOfSubarraysQueue(nums []int, k int) int {
	s := 0
	oddi := make([]int, 0, len(nums))
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			oddi = append(oddi, i)
			if len(oddi) == k+1 {
				count += ((oddi[0] - s + 1) * (i - oddi[k-1]))
				s = oddi[0] + 1
				oddi = oddi[1:]
			}
		}
		if i == len(nums)-1 && len(oddi) == k {
			count += (oddi[0] - s + 1) * (i - oddi[k-1] + 1)
		}
	}
	return count
}

func numberOfSubarraysMap(nums []int, k int) int {
	attr := make([]int, len(nums)+1)
	count := 0
	attr[0] = 0
	for i := 1; i < len(nums)+1; i++ {
		attr[i] = attr[i-1] + nums[i-1]&1
	}
	fmt.Println(attr)
	h := map[int]int{}
	for _, a := range attr {
		if c, ok := h[a]; ok {
			h[a] = c + 1
		} else {
			h[a] = 1
		}
		if c, ok := h[a-k]; ok {
			count += c
		}
	}
	fmt.Println(h)
	return count
}
