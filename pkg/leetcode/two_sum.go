package leetcode

func twoSum(nums []int, target int) []int {
	var numsMapper map[int]int
	numsMapper = make(map[int]int)
	for index, num := range nums {
		if value, ok := numsMapper[target-num]; ok {
			return []int{value, index}
		}
		if _, ok := numsMapper[num]; !ok {
			numsMapper[num] = index
		}
	}

	return []int{1, 2}
}
func twoSum2(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
	for {
		if numbers[start]+numbers[end] == target {
			return []int{start + 1, end + 1}
		} else if numbers[start]+numbers[end] < target {
			start++
		} else {
			end--
		}
		if start == end {
			break
		}
	}

	return []int{1, 2}
}
