package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if n == m {
		return head
	}
	ans := head
	var rh, rt, rp *ListNode
	for i := 1; head != nil && i <= n; i++ {
		if i == m-1 {
			rp = head
			head = head.Next
		} else if i == m {
			rt, rh, head = head, head, head.Next
			rh.Next = nil
		} else if i == n {
			head, head.Next, rh = head.Next, rh, head
			rt.Next = head
			if rp != nil {
				rp.Next = rh
			} else {
				ans = rh
			}
		} else if i > m && i < n {
			head, head.Next, rh = head.Next, rh, head
		} else {
			head = head.Next
		}
	}
	return ans
}
