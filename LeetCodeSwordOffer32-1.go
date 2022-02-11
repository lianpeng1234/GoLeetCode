package main

func levelOrder(root *TreeNode) []int {
	levelMap := make(map[int]*[]int)
	diGuiLevelOrder(root, 0, levelMap)

	var result []int
	level := 0
	for true {
		value, ok := levelMap[level]
		if ok {
			for i := 0; i < len(*value); i++ {
				xx := *value
				result = append(result, xx[i])
			}

		} else {
			break
		}
		level++
	}

	return result
}

func diGuiLevelOrder(node *TreeNode, level int, levelMap map[int]*[]int) {
	if node == nil {
		return
	}
	value, ok := levelMap[level]
	if !ok {
		value = &[]int{}
		levelMap[level] = value
	}
	*value = append(*value, node.Val)
	diGuiLevelOrder(node.Left, level+1, levelMap)
	diGuiLevelOrder(node.Right, level+1, levelMap)
}
