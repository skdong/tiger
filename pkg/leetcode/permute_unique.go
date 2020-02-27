package leetcode
import (
	"fmt"
)

func PermuteUnique(nums []int){
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	var ret [][]int
	var back_track func (first int)
	back_track = func (first int){
		if first == n-1{
			line := make([]int, 0, n)
			line = append(line, nums...)
			ret = append(ret, line)
		}else{
			for i:=first; i<n; i++{
				if i!= first && nums[i] == nums[first]{
					continue
				}
				nums[i],nums[first] = nums[first],nums[i]
				back_track(first+1)
				nums[i],nums[first] = nums[first],nums[i]
			}
		}
	}
	back_track(0)
	return ret
}