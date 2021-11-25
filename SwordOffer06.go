package main

func reversePrint(head *ListNode) []int {
	var arr []int

	for i := 0; ; i++ {
		if head == nil {
			break
		}
		arr = append(arr, head.Val)
		head = head.Next
	}

	length := len(arr)

	tmpArr := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		tmpArr[length-i-1] = arr[i]
	}

	return tmpArr
}
