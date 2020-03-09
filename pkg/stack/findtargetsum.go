package stack

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	ways := [][]int{{0, 1}}
	for i := 0; i < len(nums); i++ {
		newWays := [][]int{}
		for flag := -1; flag < 2; flag += 2 {
			cache := map[int][]int{}
			for _, way := range ways {
				newWay := []int{}
				newWay = append(newWay, way...)
				newWay = append(newWay, flag*nums[i])
				newWay[0] += flag * nums[i]
				if w, ok := cache[newWay[0]]; ok {
					w[1] += newWay[1]
					continue
				} else {
					cache[newWay[0]] = newWay
				}
				if i == len(nums)-1 && newWay[0] != S {
					continue
				}
				newWays = append(newWays, newWay)
				fmt.Println(newWays)
			}
		}
		ways = newWays
	}
	num := 0
	for _, way := range ways {
		num += way[1]
	}
	return num
}

func findTargetSumWaysD(nums []int, S int) int {
	dp := make([][2000]int, len(nums))
	sums := []int{0}
	for i := 0; i < len(nums); i++ {
		var cache []int
		newSum := []int{}
		cache = make([]int, 2000)
		for _, s := range sums {
			var lastSum int
			if i == 0 {
				lastSum = 1
			} else {
				lastSum = dp[i-1][s+1000]

			}
			for flag := -1; flag < 2; flag += 2 {
				j := s + flag*nums[i]
				if cache[1000+j] == 0 {
					newSum = append(newSum, j)
					cache[1000+j] = 1
				}
				dp[i][1000+j] += lastSum
			}
		}
		sums = newSum
	}
	return dp[len(nums)-1][S+1000]
}
