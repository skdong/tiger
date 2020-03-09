package huawei

import "fmt"

func DisBullet(nums []int) int {
	return disBullet(nums)
}

func disBullet(nums []int) int {
	count := 0
	next := make([]int, len(nums))
	for {
		overd := true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] != nums[i+1] {
				overd = false
				break
			}
		}
		if overd {
			break
		}

		for i := 0; i < len(nums); i++ {
			if nums[i]%2 == 1 {
				nums[i]++
			}
			nums[i] = nums[i] / 2
			next[(i+1)%len(nums)] = nums[i]
		}

		for i := 0; i < len(nums); i++ {
			nums[i] = nums[i] + next[i]
		}
		count++
		fmt.Println(nums)
	}
	return count
}
