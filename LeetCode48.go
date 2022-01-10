package main

func rotate(matrix [][]int) {
	length := len(matrix)
	forTimes := length / 2

	for i := 0; i < forTimes; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j], matrix[length-1-i][j] = matrix[length-1-i][j], matrix[i][j]
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
