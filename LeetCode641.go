package main

import "container/list"

type MyCircularDeque struct {
	size int
	mq   *list.List
}

func Constructor(k int) MyCircularDeque {
	mq := &list.List{}
	queen := MyCircularDeque{k, mq}
	return queen
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.size <= this.mq.Len() {
		return false
	}
	this.mq.PushFront(value)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.size <= this.mq.Len() {
		return false
	}
	this.mq.PushBack(value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.mq.Len() <= 0 {
		return false
	}
	ele := this.mq.Front()
	this.mq.Remove(ele)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.mq.Len() <= 0 {
		return false
	}
	ele := this.mq.Back()
	this.mq.Remove(ele)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.mq.Len() <= 0 {
		return -1
	}
	return this.mq.Front().Value.(int)
}

func (this *MyCircularDeque) GetRear() int {
	if this.mq.Len() <= 0 {
		return -1
	}
	return this.mq.Back().Value.(int)
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.mq.Len() <= 0 {
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if this.size == this.mq.Len() {
		return true
	}
	return false
}
