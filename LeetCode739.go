package main

import "container/list"

type stackNode struct {
	index int
	value int
}

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)

	result := [10]int{}

	stack := list.List{}

	var i int
	for i = 0; i < length; i++ {
		for true {
			ele := stack.Back()
			if ele == nil {
				break
			} else {
				stackNode := ele.Value.(*stackNode)
				if stackNode.value >= temperatures[i] {
					break
				}
				result[stackNode.index] = i - stackNode.index
				stack.Remove(ele)
			}
		}

		stack.PushBack(&stackNode{index: i, value: temperatures[i]})
	}

	for stack.Len() > 0 {
		ele := stack.Back()
		stackNode := ele.Value.(*stackNode)
		result[stackNode.index] = 0
		stack.Remove(ele)
	}

	return result[0:length]
}
