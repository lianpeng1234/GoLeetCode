package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	DiGui(root, &max, 1)
	return max
}

func DiGui(node *Node, max *int, depth int) {
	if node == nil || node.Children == nil || len(node.Children) == 0 {
		if depth > *max {
			*max = depth
		}
		return
	}

	for _, item := range node.Children {
		DiGui(item, max, depth+1)
	}
}
