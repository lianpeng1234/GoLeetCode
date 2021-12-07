package main

// 高度不一样
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var tree *TreeNode
	if root1 == nil && root2 == nil {
		return nil
	}

	tree = createNode(root1, root2)

	//
	var nextNode1 *TreeNode
	var nextNode2 *TreeNode
	if root1 != nil && root1.Left != nil {
		nextNode1 = root1.Left
	}
	if root2 != nil && root2.Left != nil {
		nextNode2 = root2.Left
	}
	digui(nextNode1, nextNode2, tree, true)

	nextNode1 = nil
	nextNode2 = nil
	if root1 != nil && root1.Right != nil {
		nextNode1 = root1.Right
	}
	if root2 != nil && root2.Right != nil {
		nextNode2 = root2.Right
	}
	digui(nextNode1, nextNode2, tree, false)

	return tree
}

func digui(node1 *TreeNode, node2 *TreeNode, tree *TreeNode, isLeft bool) {
	if node1 == nil && node2 == nil {
		return
	}
	xx := createNode(node1, node2)

	if isLeft {
		tree.Left = xx
	} else {
		tree.Right = xx
	}

	var nextNode1 *TreeNode
	var nextNode2 *TreeNode
	if node1 != nil && node1.Left != nil {
		nextNode1 = node1.Left
	}
	if node2 != nil && node2.Left != nil {
		nextNode2 = node2.Left
	}
	digui(nextNode1, nextNode2, xx, true)

	nextNode1 = nil
	nextNode2 = nil
	if node1 != nil && node1.Right != nil {
		nextNode1 = node1.Right
	}
	if node2 != nil && node2.Right != nil {
		nextNode2 = node2.Right
	}
	digui(nextNode1, nextNode2, xx, false)
}

func createNode(node1 *TreeNode, node2 *TreeNode) *TreeNode {
	val1 := 0
	if node1 != nil {
		val1 = node1.Val
	}
	val2 := 0
	if node2 != nil {
		val2 = node2.Val
	}

	return &TreeNode{val1 + val2, nil, nil}
}
