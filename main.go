package main

func main() {
	node9 := &TreeNode{9, nil, nil}
	node20 := &TreeNode{20, nil, nil}
	node5 := &TreeNode{5, nil, nil}

	node10 := &TreeNode{10, node9, node20}
	node8 := &TreeNode{8, node5, node10}

	convertBST(node8)
}
