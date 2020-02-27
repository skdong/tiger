package leetcode

func maxArea(height []int)int{
	var first, last, max int
	max = 0
	first = 0
	last = len(height)
	for first != last{
		if height[first] > height[last]{
			if max < height[last]*(last - first){
				max = height[last]*(last-first)
			}
			first ++
		}else{
			if max < height[last]*(last - first){
				max = height[last]*(last-first)
			}
			last  --
		}
	}
	return max

}