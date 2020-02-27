package leetcode

func strongPasswordChecker(s string) int {
	if len(s) < 6 || len(s) > 20{
		return -1
	}
	
	uperFlag := false
	lowFlag := false
	digitFlag := false

	count := 0
	allCounts := []int{}
	var last uint8
	for i := range s{
		if s[i] >= 'A' && s[i] <= 'Z'{
			uperFlag = true
		}else if s[i] >= 'a' && s[i] <= 'z'{
			lowFlag = true
		}else if s[i] >= '0' && s[i] <= '9'{
			digitFlag = true
		}

		if count == 0{
			last = s[i]
		}else if last == s[i]{
			count ++
		}else if last != s[i]{
			if count > 2{
				allCounts = append(allCounts, count)
			}
			last = s[i]
			count = 1
		}

	}

	if uperFlag && lowFlag && digitFlag && len(allCounts) == 0{
		return 0
	}else{
		return -1
	}
	
}