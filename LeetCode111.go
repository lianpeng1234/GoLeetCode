package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	treeLevel := make(map[int][]*TreeNode)

	recursion(root, 1, &treeLevel)

	for i := 1; ; i++ {
		level := treeLevel[i]
		if level == nil || len(level) < 1 {
			break
		}
		for _, item := range level {
			if item.Right == nil && item.Left == nil {
				return i
			}
		}
	}

	return 0
}

func recursion(node *TreeNode, depth int, treeLevel *map[int][]*TreeNode) {
	if node == nil {
		return
	}
	level, ok := (*treeLevel)[depth]
	if ok {
		level = append(level, node)
		(*treeLevel)[depth] = level
	} else {
		var tmp []*TreeNode
		tmp = append(tmp, node)
		(*treeLevel)[depth] = tmp
	}
	recursion(node.Left, depth+1, treeLevel)
	recursion(node.Right, depth+1, treeLevel)
}
