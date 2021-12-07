package main

func minArray(numbers []int) int {
	if numbers == nil || len(numbers) == 0 {
		return 0
	}
	length := len(numbers) - 1
	for i := 0; i < length; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}
