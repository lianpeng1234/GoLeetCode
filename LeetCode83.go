package main

func deleteDuplicatesXX(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tmpHead := head

	for i := 0; ; i++ {
		if tmpHead.Next == nil {
			break
		}
		if tmpHead.Val == tmpHead.Next.Val {
			tmpHead.Next = tmpHead.Next.Next
		} else {
			tmpHead = tmpHead.Next
		}
	}

	return head
}
