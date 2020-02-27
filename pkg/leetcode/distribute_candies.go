package leetcode


func distributeCandies(candies int, num_people int) []int {
	var ans []int
	ans = make([]int, num_people)
	sum := candies
	for i:=0; sum >0; i++{
		if i+1 < sum{
			ans[i%num_people] += (i+1)
		}else{
			ans[i%num_people] += sum
		}
		sum -= (i+1)
	}
	return ans
}