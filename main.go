package main

import "fmt"

func main() {

//["MyCircularDeque","insertFront","deleteLast","getRear","getFront","getFront","deleteFront","insertFront","insertLast","insertFront","getFront","insertFront"]
//	[[4],                [9],            [],        [],       [],         [],           [],       [6],           [5],         [9],         [],        [6]]

	circularDeque := Constructor(4)           // 设置容量大小为3
	fmt.Println(circularDeque.InsertFront(5))  // 返回 true
	fmt.Println(circularDeque.InsertLast(7))  // 返回 true
	fmt.Println(circularDeque.GetFront()) // 返回 true
	fmt.Println(circularDeque.InsertFront(4)) // 已经满了，返回 false
	fmt.Println(circularDeque.GetRear())      // 返回 2
	fmt.Println(circularDeque.IsFull())       // 返回 true
	fmt.Println(circularDeque.DeleteLast())   // 返回 true
	fmt.Println(circularDeque.InsertFront(4)) // 返回 true
	fmt.Println(circularDeque.GetFront())     // 返回 4
}
