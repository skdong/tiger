package huawei

func corpFlightBookings(bookings [][]int, n int) []int {
	/*answer 记录节点增量，开始加，结束减*/
	answer := make([]int, n)
	for _, b := range bookings {
		answer[b[0]-1] += b[2]
		if b[1] < n {
			answer[b[1]] -= b[2]
		}

	}
	for i := 1; i < n; i++ {
		answer[i] += answer[i-1]
	}
	return answer
}
