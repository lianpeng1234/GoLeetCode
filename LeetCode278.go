package main

func firstBadVersion(n int) int {
	// 二分查找变体，查找第一个等于目标值的元素
	low := 0
	height := n
	center := 0
	for low <= height {
		center = (height-low)/2 + low
		if isBadVersion(center) {
			if center >= 1 && isBadVersion(center-1) {
				height = center - 1
			} else {
				return center
			}
		} else {
			low = center + 1
		}
	}
	return center
}

func isBadVersion(version int) bool {
	return false
}
