package leetcode

import (
	"fmt"
)

func CombinationSum(candidates []int, target int) {
	fmt.Println( combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	var back_track func (first int, cur_target int, stack []int)
	var out_put [][]int
	back_track = func(first int, cur_target int, stack[]int){
		n := len(candidates)
		if cur_target == 0{
			line := make([]int, 0, len(stack))
			line = append(line, stack...)
			out_put = append(out_put, line)
		}else if cur_target >0{
			for i:=first; i<n; i++{
				back_track(i, cur_target-candidates[i], append(stack,candidates[i]))
			}
		}else{
				return
		}
	}
	var first_stack []int
	back_track(0, target, first_stack)
	return out_put
}