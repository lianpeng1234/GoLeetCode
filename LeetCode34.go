package main

// 二分变体，查找第一个等于目标值的索引下标
func searchRange(nums []int, target int) []int {
	if nums == nil {
		return []int{-1, -1}
	}
	low := 0
	height := len(nums) - 1

	center := 0
	for low <= height {
		center = (height-low)/2 + low
		if nums[center] == target {
			if center >= 1 && nums[center-1] == target {
				height = center - 1
			} else {
				break
			}
		} else if nums[center] < target {
			low = center + 1
		} else {
			height = center - 1
		}
	}

	position := []int{}
	for i := center; i < len(nums); i++ {
		if nums[i] == target {
			position = append(position, i)
			continue
		}
		break
	}

	if len(position) < 1 {
		return []int{-1, -1}
	} else if len(position) == 1{
		return []int{position[0], position[0]}
	} else {
		return []int{position[0], position[len(position)-1]}
	}
}
