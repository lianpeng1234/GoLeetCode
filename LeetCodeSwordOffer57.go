package main

func twoSum(nums []int, target int) []int {
	cmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cmap[nums[i]] = nums[i]
	}

	var result []int
	for i := 0; i < len(nums); i++ {
		a := target - nums[i]
		value, ok := cmap[a]
		if ok {
			result = append(result, nums[i])
			result = append(result, value)
			return result
		}
	}
	return result
}
