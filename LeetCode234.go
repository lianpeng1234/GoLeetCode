package main

import "container/list"

func IsPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	stack := list.List{}
	for i := 0; ; i++ {
		if head == nil {
			break
		}
		stack.PushBack(head)
		head = head.Next
	}

	for i := 0; ; i++ {
		if stack.Len() <= 1 {
			return true
		}
		front := stack.Front()
		before := front.Value.(*ListNode)

		back := stack.Back()
		end := back.Value.(*ListNode)
		if before.Val == end.Val {
			stack.Remove(front)
			stack.Remove(back)
		} else {
			return false
		}
	}
}
