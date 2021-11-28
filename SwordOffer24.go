package main

func reverseList(head *ListNode) *ListNode {
	var tmp *ListNode

	var pre *ListNode
	pre = head

	var cur *ListNode

	for i := 0; ; i++ {
		if pre == nil {
			break
		}
		tmp = pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
	}

	return cur
}
