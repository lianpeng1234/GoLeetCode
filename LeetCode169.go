package main

func majorityElement(nums []int) int {
	length := len(nums)
	maps := make(map[int]int)
	for i := 0; i < length; i++ {
		value, _ := maps[nums[i]]
		value++
		maps[nums[i]] = value
		if value > length/2 {
			return nums[i]
		}
	}
	return -1
}
