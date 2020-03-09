package dp

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for _, c := range coins {
		for a := c; a <= amount; a++ {
			if a > c && dp[a-c] != 0 {
				if dp[a] == 0 || dp[a-c]+1 < dp[a] {
					dp[a] = dp[a-c] + 1
				}
			} else if a == c {
				dp[a] = 1
				continue
			}
		}
	}
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

func coinchangeBack(coins []int, amount int) int {
	return coinchangeBackN(0, coins, amount)
}

func coinchangeBackN(n int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	ans := -1
	if n < len(coins) {
		maxVal := amount / coins[n]
		for i := 0; i <= maxVal; i++ {
			n1min := coinchangeBackN(n+1, coins, amount-i*coins[n])
			if n1min != -1 && (ans == -1 || n1min+i < ans) {
				ans = n1min + i
			}
		}

	}
	return ans
}

func coinChangeDp(coins []int, amount int) int {
	cache := map[int]int{}
	return coinChangeDpC(cache, coins, amount)
}
func coinChangeDpC(cache map[int]int, coins []int, amount int) int {
	if v, ok := cache[amount]; ok {
		return v
	}
	if amount == 0 {
		return 0
	}
	ans := -1
	for _, c := range coins {
		ansL := coinChangeDp(coins, amount-c)
		if ansL != -1 && (ans == -1 || ansL+1 < ans) {
			ans = ansL + 1
			cache[amount] = ans
		}
	}
	return ans
}
