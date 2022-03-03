package main

func getIntersectionNodeX(headA, headB *ListNode) *ListNode {
	tmpA := headA
	aLen := getListLength(tmpA)

	tmpB := headB
	bLen := getListLength(tmpB)

	tmpA = headA
	tmpB = headB
	if aLen > bLen {
		skip := 0
		tmpA = headA
		for tmpA != nil {
			skip++
			tmpA = tmpA.Next
			if skip == aLen-bLen {
				break
			}
		}
	} else if bLen > aLen {
		skip := 0
		for tmpB != nil {
			skip++
			tmpB = tmpB.Next
			if skip == bLen-aLen {
				break
			}
		}
	}

	for tmpA != nil {
		if tmpA == tmpB {
			break
		}
		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}

	return tmpA
}

func getListLength(head *ListNode) int {
	a := 0
	for head != nil {
		a++
		head = head.Next
	}
	return a
}
