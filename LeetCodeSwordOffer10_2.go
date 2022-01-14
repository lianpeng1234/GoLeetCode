package main

func numWays(n int) int {
	result := make(map[int]int)
	return diGuiNumWays(n, result)
}

func diGuiNumWays(n int, result map[int]int) int {
	if n < 2 {
		result[n] = 1
		return 1
	}

	a := 0
	if value, ok := result[n-1]; ok {
		a = value
	}
	if a == 0 {
		a = diGuiNumWays(n-1, result)
	}

	b := 0
	if value, ok := result[n-2]; ok {
		b = value
	}
	if b == 0 {
		b = diGuiNumWays(n-2, result)
	}

	sum := (a + b) % 1000000007
	result[n] = sum
	return sum
}
