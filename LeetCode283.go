package main

func moveZeroes(nums []int) {
	length := len(nums)
	slow := 0
	quick := 0
	for quick < length && slow < length {
		if nums[slow] == 0 && nums[quick] != 0 {
			nums[slow], nums[quick] = nums[quick], nums[slow]
			slow++
			quick++
		} else if nums[slow] != 0 {
			slow++
			quick = slow + 1
		} else if nums[quick] == 0 {
			quick++
		}
	}
}
