package main

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val && head.Next == nil {
		return nil
	}
	if head.Val == val && head.Next != nil {
		return head.Next
	}

	tmpHead := head

	var tmpNode *ListNode
	for tmpHead != nil {
		if tmpHead.Val == val {
			tmpNode.Next = tmpHead.Next
			break
		} else {
			tmpNode = tmpHead
			tmpHead = tmpHead.Next
		}

	}

	return head
}
