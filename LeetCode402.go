package main

import "strings"

func RemoveKdigits(num string, k int) string {
	stack := []byte{}
	stack = append(stack, num[0])
	for i := 1; i < len(num); i++ {
		item := num[i]
		for xx := 0; ; xx++ {
			if len(stack) > 0 && k > 0 {
				if item < stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
					k--
				} else {
					break
				}
			} else {
				break
			}
		}
		stack = append(stack, item)
	}

	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
