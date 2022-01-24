package main

import "fmt"

func main() {
	listNode1 := &ListNode{1, nil}
	listNode2 := &ListNode{2, listNode1}
	listNode3 := &ListNode{3, listNode2}
	listNode4 := &ListNode{4, listNode3}
	xx := deleteNode(listNode4, 1)
	fmt.Println(xx)
}
