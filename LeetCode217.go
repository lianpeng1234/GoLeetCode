package main

// 1  2  6  7  4
// 1  2  4  7  6
func containsDuplicate(nums []int) bool {
	QuickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func QuickSort(nums []int, start int, end int) []int {
	if start >= end {
		return nums
	}
	ponit := nums[end]
	i := start
	for j := i; j < end; j++ {
		if nums[j] < ponit {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]

	QuickSort(nums, start, i-1)
	QuickSort(nums, i+1, end)

	return nums
}
