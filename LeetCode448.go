package main

func findDisappearedNumbers(nums []int) []int {
	startIndex := 0
	length := len(nums)
	for startIndex < length {
		if nums[startIndex] == startIndex+1 {
			startIndex++
			continue
		}

		target := nums[startIndex] - 1
		if nums[target] == target+1 {
			startIndex++
			continue
		}

		nums[startIndex], nums[target] = nums[target], nums[startIndex]
	}

	var result []int
	for i := 0; i < length; i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}

	return result
}
