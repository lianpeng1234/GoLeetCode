package main

//1，先纵向查找第一个大于等于目标值的元素下标，记为 x
//2，在 matrix[0] 到 matrix[x] 之间查找目标值是否存在
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	x := binarySearch1(matrix, target)
	for i := x; i >= 0; i-- {
		targetIndex := binarySearch2(matrix[i], target)
		if targetIndex != -1 {
			return true
		}
	}
	return false
}

// 先纵向查找第一个大于等于目标值的元素下标，记为 x
func binarySearch1(matrix [][]int, target int) int {
	low := 0
	height := len(matrix) - 1
	for height >= low {
		center := (height-low)/2 + low
		if len(matrix[center]) < 1 {
			low++
			continue
		}
		if matrix[center][0] >= target {
			if center >= 1 && len(matrix[center-1]) > 0 && matrix[center-1][0] >= target {
				height = center - 1
			} else {
				return center
			}
		} else {
			low = center + 1
		}
	}
	return len(matrix) - 1
}

// 在 matrix[0] 到 matrix[x] 之间查找目标值是否存在
func binarySearch2(nums []int, target int) int {
	low := 0
	hight := len(nums) - 1
	for hight >= low {
		center := (hight-low)/2 + low
		if nums[center] == target {
			return center
		} else if nums[center] < target {
			low = center + 1
		} else {
			hight = center - 1
		}
	}
	return -1
}
