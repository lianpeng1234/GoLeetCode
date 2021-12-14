package main

// 先把第一个链表、第二个链表合并为一个，得到新链表X，再把X与第三方链表进行合并，依次类推
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	length := len(lists)
	if length <=0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	tmpList := mergeTowLists(lists[0], lists[1])

	for i := 2; i < len(lists); i++ {
		tmpList = mergeTowLists(tmpList, lists[i])
	}

	return tmpList
}

func mergeTowLists(one *ListNode, two *ListNode) *ListNode {
	head := &ListNode{}
	tmpHead := head
	for one != nil && two != nil {
		if one.Val <= two.Val {
			tmpHead.Next = one
			one = one.Next
			tmpHead = tmpHead.Next
		} else {
			tmpHead.Next = two
			two = two.Next
			tmpHead = tmpHead.Next
		}
	}

	if one == nil && two != nil {
		tmpHead.Next = two
	} else if one != nil && two == nil {
		tmpHead.Next = one
	}

	return head.Next
}
