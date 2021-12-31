package main

// 异或运算，相同得 0，不同的 1
// a ^ 0 = a
// a ^ a = 0
// a ^ b ^ a = a ^ a ^ b = 0 ^ b = b (异或操作满足交换律)
func singleNumber(nums []int) int {
	single := 0
	for i := 0; i < len(nums); i++ {
		single = single ^ nums[i]
	}
	return single
}
