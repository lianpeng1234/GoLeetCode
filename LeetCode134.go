package main

// 加油站问题，超过时间限制
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	for i := 0; i < length; i++ {
		current := i
		run := false
		surplus := 0 // 油箱油量
		for true {
			if run && current == i {
				return i
			}
			if surplus+gas[current] < cost[current] {
				i = current
				break
			} else {
				run = true
				surplus = surplus + gas[current] - cost[current]
				if current == length-1 {
					current = 0
				} else {
					current++
				}
			}
		}
	}
	return -1
}
