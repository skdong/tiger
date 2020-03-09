package huawei

import "fmt"

func quickSort(nums []int) {
	s_i := 0
	e_i := len(nums) - 1
	for s_i < e_i {
		if nums[s_i] <= nums[s_i+1] {
			nums[s_i], nums[s_i+1] = nums[s_i+1], nums[s_i]
			s_i++
		} else if nums[s_i] > nums[e_i] {
			e_i--
		} else {
			nums[s_i], nums[s_i+1], nums[e_i] = nums[e_i], nums[s_i], nums[s_i+1]
			s_i++
		}
	}
	if len(nums) > 2 {
		quickSort(nums[s_i+1:])
		quickSort(nums[:s_i])
	}
}

func JobToPrint(jobs []int, index int) (int, []int) {
	times := 0
	q := make([][2]int, len(jobs))
	out := make([]int, len(jobs))
	order := make([]int, len(jobs))
	for i, p := range jobs {
		q[i] = [2]int{i, p}
		order[i] = p
	}
	quickSort(order)
	fmt.Println(order)
	max_i := 0
	for len(q) > 0 {
		fmt.Println(q)
		i, p := q[0][0], q[0][1]
		if p == order[max_i] {
			q = q[1:]
			out[max_i] = i
			if i == index {
				times = max_i + 1
			}
			max_i++
		} else {
			q = append(q[1:], q[0])
		}
	}
	return times, out
}
