package main

func RemoveElement(nums []int, val int) int {
	length := len(nums)
	endIndex := length - 1
	for i := 0; ; i++ {
		if i >= length {
			break
		}
		if nums[i] < 0 {
			break
		}
		if nums[i] == val {
			nums[i] = nums[endIndex]
			nums[endIndex] = -1
			endIndex--
			i--
		}
	}
	return endIndex + 1
}
