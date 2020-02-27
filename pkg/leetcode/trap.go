package leetcode

func trap(height []int)int{
	var waters,  left, right, left_acc, right_acc int
	waters = 0
	n := len(height)
	left = 0
	left_acc = 0
	right = n - 1
	right_acc = 0
	for i:=1; i<n ; i++{
		if left == right{
			break
		}
		if height[i]>= height[left]{
			waters += left_acc
			left_acc = 0
			left = i
		}else{
			left_acc += (height[left] - height[i])
		}

		if height[n-i-1]>= height[right]{
			waters += right_acc
			right_acc = 0
			right = n -i -1
		}else{
			right_acc += (height[right] - height[n-1-i])
		}
	}
	return waters
}