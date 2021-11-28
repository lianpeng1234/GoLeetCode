package main

func main() {
	fore := NewListNode(4, nil)
	three := NewListNode(3, fore)
	two := NewListNode(2, three)
	one := NewListNode(1, two)

	ReverseBetween(one, 1, 2)
}
