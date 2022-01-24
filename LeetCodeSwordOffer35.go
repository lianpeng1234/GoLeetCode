package main

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cmap := make(map[*Node]int)

	count := 0
	tmpHead := head
	for tmpHead != nil {
		cmap[tmpHead] = count
		tmpHead = tmpHead.Next
		count++
	}

	var location []int
	tmpHead = head
	for tmpHead != nil {
		value, ok := cmap[tmpHead.Random]
		if ok {
			location = append(location, value)
		} else {
			location = append(location, -1)
		}
		tmpHead = tmpHead.Next
	}

	// 开始 copy
	tmpHead = head
	copyLinkHead := copyLink(tmpHead)
	//
	cmap2 := make(map[int]*Node)
	count = 0
	tmpHead = copyLinkHead
	for tmpHead != nil {
		cmap2[count] = tmpHead
		count++
		tmpHead = tmpHead.Next
	}
	//
	tmpHead = copyLinkHead
	count = 0
	for tmpHead != nil {
		if location[count] > -1 {
			tmpHead.Random = cmap2[location[count]]
		} else {
			tmpHead.Random = nil
		}
		tmpHead = tmpHead.Next
		count++
	}

	return copyLinkHead
}

func copyLink(head *Node) *Node {
	if head.Next == nil {
		node := &Node{Val: head.Val, Next: nil, Random: nil}
		return node
	}
	next := copyLink(head.Next)
	node := &Node{Val: head.Val, Next: next, Random: nil}
	return node
}
