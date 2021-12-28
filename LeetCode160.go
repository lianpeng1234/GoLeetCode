package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA := 0
	tmpHeadA := headA
	for true {
		if tmpHeadA == nil {
			break
		}
		lengthA++
		tmpHeadA = tmpHeadA.Next
	}

	lengthB := 0
	tmpHeadB := headB
	for true {
		if tmpHeadB == nil {
			break
		}
		lengthB++
		tmpHeadB = tmpHeadB.Next
	}

	cha := 0
	if lengthA > lengthB {
		cha = lengthA - lengthB
	} else {
		cha = lengthB - lengthA
	}

	tmpHeadA = headA
	tmpHeadB = headB
	if lengthA > lengthB {
		for i := 0; i < cha; i++ {
			tmpHeadA = tmpHeadA.Next
		}
		for true {
			if tmpHeadA == nil {
				return nil
			} else {
				if tmpHeadA.Val == tmpHeadB.Val {
					return tmpHeadA
				}
				tmpHeadA = tmpHeadA.Next
				tmpHeadB = tmpHeadB.Next
			}
		}
	} else {
		for i := 0; i < cha; i++ {
			tmpHeadB = tmpHeadB.Next
		}
		for true {
			if tmpHeadA == nil {
				return nil
			} else {
				if tmpHeadA == tmpHeadB {
					return tmpHeadA
				}
				tmpHeadA = tmpHeadA.Next
				tmpHeadB = tmpHeadB.Next
			}
		}
	}
	return nil
}
