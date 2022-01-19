package main

func isSymmetric(root *TreeNode) bool {
	var list1 []int
	// 前序遍历
	preorder(root, &list1, 0)

	// 镜像处理
	mirrorImage(root)

	var list2 []int
	// 前序遍历
	preorder(root, &list2, 0)

	for i := 0; i <len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// 先续遍历 location 0：root节点，1：坐节点，2：右节点
func preorder(node *TreeNode, list *[]int, location int) {
	if node == nil {
		return
	}
	if location == 0 {
		*list = append(*list, node.Val)
	} else if location == 1 {
		*list = append(*list, node.Val+1000000)
	} else if location == 2 {
		*list = append(*list, node.Val+2000000)
	}
	preorder(node.Left, list, 1)
	preorder(node.Right, list, 2)
}

// 镜像处理
func mirrorImage(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	mirrorImage(node.Left)
	mirrorImage(node.Right)
}
