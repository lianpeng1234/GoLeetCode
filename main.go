package main

import "fmt"

func main() {
	preorder := []int{1, 2, 3}
	inorder := []int{3, 2, 1}
	node := buildTree(preorder, inorder)
	fmt.Println(node)
}
