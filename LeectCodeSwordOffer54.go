package main

func kthLargest(root *TreeNode, k int) int {
	var nums []int
	recursionKthLargest(root, &nums)
	return nums[len(nums)-k]
}

// 中序遍历
func recursionKthLargest(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	recursionKthLargest(node.Left, nums)
	*nums = append(*nums, node.Val)
	recursionKthLargest(node.Right, nums)
}
