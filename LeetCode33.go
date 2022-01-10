package main

func search(nums []int, target int) int {
	// 搜索旋转的起始下标
	k := searchK(nums)

	low := 0
	height := len(nums) - 1

	if nums[k] == target {
		return k
	} else if nums[k] <= target && target <= nums[height] {
		low = k
	} else {
		height = k - 1
	}

	for low <= height {
		center := (height-low)/2 + low
		if nums[center] == target {
			return center
		} else if nums[center] < target {
			low = center + 1
		} else {
			height = center - 1
		}
	}

	return -1
}

// 搜索旋转的起始下标
func searchK(nums []int) int {
	k := 0
	low := 0
	height := len(nums) - 1
	for low <= height {
		center := (height-low)/2 + low

		if center < len(nums)-1 && nums[center] > nums[center+1] {
			k = center + 1
			break
		}

		if center >= 1 && nums[center-1] > nums[center] {
			k = center
			break
		}

		if nums[low] < nums[center] {
			low = center + 1
		} else {
			height = center - 1
		}
	}

	return k
}