package main

func maxDepth_lp(root *TreeNode) int {
	depth := -1
	recursionMaxDepth(root, &depth, 1)
	return depth
}

func recursionMaxDepth(node *TreeNode, depth *int, level int) {
	if node == nil {
		if *depth < level-1 {
			*depth = level - 1
		}
		return
	}
	recursionMaxDepth(node.Left, depth, level+1)
	recursionMaxDepth(node.Right, depth, level+1)
}
