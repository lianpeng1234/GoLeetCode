package main

import "fmt"

func main() {
	// [10,12,6,8,3,11]
	//[10,12,6,8]

	t11 := &TreeNode{11, nil, nil}
	t8 := &TreeNode{8, nil, nil}
	t3 := &TreeNode{3, nil, nil}

	t6 := &TreeNode{6, t11, nil}
	t12 := &TreeNode{12, t8, t3}
	t10 := &TreeNode{10, t12, t6}

	t6b := &TreeNode{6, nil, nil}
	t12b := &TreeNode{12, t8, nil}
	t10b := &TreeNode{10, t12b, t6b}


	xx := isSubStructure(t10, t10b)
	fmt.Println(xx)
}
