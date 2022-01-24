package main

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) < 1{
		return nil
	}

	var result []int

	columnStartIndex := 0
	columnEndIndex := len(matrix[0]) - 1

	rowStartIndex := 0
	rowEndIndex := len(matrix) - 1

	allCount := len(matrix[0]) * len(matrix)
	count := 0

	// 方向 1 右，2 左，3 下，4 上
	direction := 1

	for true {
		// --》
		if direction == 1 {
			for i := columnStartIndex; i <= columnEndIndex; i++ {
				result = append(result, matrix[rowStartIndex][i])
				count++
			}
			rowStartIndex++
		}
		if count >= allCount {
			break
		}
		direction = 3
		// ↓
		if direction == 3 {
			for i := rowStartIndex; i <= rowEndIndex; i++ {
				result = append(result, matrix[i][columnEndIndex])
				count++
			}
			columnEndIndex--
		}
		if count >= allCount {
			break
		}
		direction = 2
		// 《--
		if direction == 2 {
			for i := columnEndIndex; i >= columnStartIndex; i-- {
				result = append(result, matrix[rowEndIndex][i])
				count++
			}
			rowEndIndex--
		}
		if count >= allCount {
			break
		}
		direction = 4
		// ↑
		if direction == 4 {
			for i := rowEndIndex; i >= rowStartIndex; i-- {
				result = append(result, matrix[i][columnStartIndex])
				count++
			}
			columnStartIndex++
		}
		if count >= allCount {
			break
		}
		direction = 1
	}

	return result
}
