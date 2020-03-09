package dp

/*
https://leetcode-cn.com/problems/decode-ways/
*/
func numDecodings(s string) int {
	dp1, dp2 := 1, 1
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] == '0' {
				return 0
			}
			continue
		}
		if s[i-1] == '0' {
			if s[i] == '0' {
				return 0
			}
			dp2 = dp1
		} else if s[i] == '0' {
			if s[i-1] > '2' {
				return 0
			}
			dp1, dp2 = dp2, dp1
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] < '7' {
			dp1, dp2 = dp1+dp2, dp1
		} else {
			dp2 = dp1
		}
	}
	return dp1
}
