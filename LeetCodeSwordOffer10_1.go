package main

func fib(n int) int {
	resultMap := make(map[int]int)
	return diGuiFib(n, resultMap)
}

func diGuiFib(n int, resultMap map[int]int) int {
	if n < 2 {
		resultMap[n] = n
		return n
	}

	a := -1
	if value, ok := resultMap[n-1]; ok {
		a = value
	}
	if a == -1 {
		a = diGuiFib(n-1, resultMap)
	}

	b := -1
	if value, ok := resultMap[n-2]; ok {
		b = value
	}
	if b == -1 {
		b = diGuiFib(n-2, resultMap)
	}

	sum := a + b
	sum = sum % 1000000007
	resultMap[n] = sum
	return sum
}
