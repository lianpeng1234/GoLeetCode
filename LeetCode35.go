package main

// 二分查找变体，找到第一个大于等于目标值的元素
func searchInsert(nums []int, target int) int {
	index := find(nums, target)
	if index != -1 {
		return index
	}

	return len(nums)
}

func find(nums []int, target int) int {
	low := 0
	height := len(nums) - 1
	center := 0
	for low <= height {
		center = (height-low)/2 + low
		if nums[center] >= target {
			if center >= 1 && nums[center-1] >= target {
				height = center - 1
			} else {
				return center
			}
		} else {
			low = center + 1
		}
	}
	return -1
}
