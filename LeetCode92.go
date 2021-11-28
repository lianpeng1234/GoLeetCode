package main

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	mute := ListNode{Val: -1, Next: head}

	l := 0
	r := 0

	var before *ListNode
	var bb *ListNode
	bb = &mute

	for i := 0; ; i++ {
		before = head
		l = l + 1
		if l == left {
			break
		}
		head = head.Next
		bb = before
	}

	var tmp *ListNode
	var cur *ListNode
	pre := before
	for i := 0; ; i++ {
		if r > right-left {
			break
		}
		r = r + 1

		tmp = pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
	}

	bb.Next = cur
	before.Next = tmp

	if left == 1 {
		before.Next = tmp
		return cur
	}

	return mute.Next
}
