package main

/**
preorder 前序遍历
inorder  中序遍历
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	inmap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inmap[inorder[i]] = i
	}

	pre := 0
	node := recursionBuildTree(preorder, pre, inorder, 0, len(inorder)-1, inmap)

	return node
}

func recursionBuildTree(preorder []int, pre int, inorder []int, inBegin int, inEnd int, inmap map[int]int) *TreeNode {
	if inBegin > inEnd {
		return nil
	}

	inorderIndex := inmap[preorder[pre]]

	rootNode := &TreeNode{Val: inorder[inorderIndex], Left: nil, Right: nil}
	rootNode.Left = recursionBuildTree(preorder, pre+1, inorder, inBegin, inorderIndex-1, inmap)
	rootNode.Right = recursionBuildTree(preorder, pre+inorderIndex-inBegin+1, inorder, inorderIndex+1, inEnd, inmap)
	return rootNode
}
