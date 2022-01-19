package main

type Result struct {
	result       bool
	count        int
	currentCount int
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	result := &Result{false, 0, 0}

	// 获取 tree 的节点数量
	getTreeNodeCount(B, result)

	diGuiIsSubStructure(A, B, result)

	return result.result
}

func diGuiIsSubStructure(father *TreeNode, sun *TreeNode, result *Result) {
	if result.result {
		return
	}
	if father == nil || sun == nil {
		return
	}
	match(father, sun, 0, result)
	if !result.result {
		result.currentCount = 0
	}
	diGuiIsSubStructure(father.Left, sun, result)
	diGuiIsSubStructure(father.Right, sun, result)
}

func match(father *TreeNode, sun *TreeNode, currentCount int, result *Result) {
	if result.count == result.currentCount {
		result.result = true
		return
	}
	if father == nil || sun == nil {
		return
	}
	if father.Val != sun.Val {
		return
	}
	result.currentCount = result.currentCount + 1
	match(father.Left, sun.Left, result.currentCount, result)
	match(father.Right, sun.Right, result.currentCount, result)
}

// 获取 tree 的节点数量
func getTreeNodeCount(node *TreeNode, result *Result) {
	if node == nil {
		return
	}
	result.count = result.count + 1
	getTreeNodeCount(node.Left, result)
	getTreeNodeCount(node.Right, result)
}
