package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow := head
	fast := head

	for i := 1; ; i++ {
		if fast == nil {
			break
		}
		fast = fast.Next
		if i > k {
			slow = slow.Next
		}
	}
	return slow
}
