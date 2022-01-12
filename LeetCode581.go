package main

// 10,2,3,4,5,10
// 7,8,9,10,2,3,4,5,10
// 7,8,9,1,2,4
// 1,1,1,1,1,1
// 1,2,3,10,9,11,12,13
func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	low := 0
	for true {
		if low+1 < length && nums[low] <= nums[low+1] {
			low++
		} else {
			break
		}
	}

	height := length - 1
	for true {
		if height-1 > -1 && nums[height-1] <= nums[height] {
			height--
		} else {
			break
		}
	}

	if height-low <= 0 {
		return 0
	}

	// 在 low 到 height 之间 是无序的  [low,height]
	// 找到 [low, height] 中的最大值，最小值
	min := 123123123
	max := -123123123
	for i := low; i <= height; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	low--
	height++

	for true {
		if low < 0 {
			break
		}
		if nums[low] <= min {
			break
		} else {
			low--
		}
	}
	for true {
		if height >= length {
			break
		}
		if nums[height] >= max {
			break
		} else {
			height++
		}
	}

	low++
	height--

	return height - low + 1
}
