package main

func findContinuousSequence(target int) [][]int {
	var result [][]int

	sum := 0
	start := 1
	end := target/2 + 1
	for start <= end {
		var xx []int
		for i := start; i <= end; i++ {
			sum = sum + i
			xx = append(xx, i)
			if sum == target {
				result = append(result, xx)
			} else if sum > target {
				break
			}
		}
		sum = 0
		start++
	}

	return result
}
