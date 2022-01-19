package main

// 双指针，夹逼法
func exchange(nums []int) []int {
	a := 0
	b := len(nums) - 1
	for a < b {
		// 交换
		if nums[a]%2 == 0 && nums[b]%2 == 1 {
			nums[a], nums[b] = nums[b], nums[a]
			a++
			b--
			continue
		}
		// 寻找偶数
		if nums[a]%2 == 1 {
			a++
		}
		// 寻找奇数
		if nums[b]%2 == 0 {
			b--
		}
	}
	return nums
}
