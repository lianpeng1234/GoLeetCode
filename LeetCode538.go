package main

func convertBST(root *TreeNode) *TreeNode {
	//mark := make(map[*TreeNode]int8)
	max := 0
	diGui(root, &max)

	return root
}

func diGui(node *TreeNode, max *int) {
	if node == nil {
		return
	}
	if node.Right != nil {
		diGui(node.Right, max)
	}

	*max = *max + node.Val
	node.Val = *max

	if node.Left != nil {
		diGui(node.Left, max)
	}
}
