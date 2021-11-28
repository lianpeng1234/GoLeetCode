package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	mute := &ListNode{-10000, head}

	tmpHead := mute

	preHead := mute

	change := false

	for i := 0; ; i++ {
		if head.Next == nil {
			if change {
				preHead.Next = nil
				change = false
			}
			break
		}
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			change = true
		} else {
			if !change {
				preHead = head
			}
			head = head.Next
			if change {
				preHead.Next = head
				change = false
			}
		}
	}

	return tmpHead.Next
}
