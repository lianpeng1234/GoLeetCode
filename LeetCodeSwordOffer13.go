package main

type Lc struct {
	row    int
	column int
	k      int
	count  int
}

func movingCount(m int, n int, k int) int {
	var lc Lc
	mark := [100][100]int{}
	lc = Lc{m, n, k, 0}
	diGuiMovingCount(0, 0, &lc, &mark)
	return lc.count
}

func diGuiMovingCount(row int, column int, lc *Lc, mark *[100][100]int) {
	if lc.row-1 < row || row < 0 || lc.column-1 < column || column < 0 {
		return
	}
	if lc.k < getSum(row, column) {
		return
	}
	if mark[row][column] == 1 {
		return
	}

	mark[row][column] = 1
	lc.count += 1
	diGuiMovingCount(row-1, column, lc, mark)
	diGuiMovingCount(row+1, column, lc, mark)
	diGuiMovingCount(row, column+1, lc, mark)
	diGuiMovingCount(row, column-1, lc, mark)
}

func getSum(row int, column int) int {
	sum := 0
	sum = sum + row%10
	sum = sum + row/10
	sum = sum + column%10
	sum = sum + column/10
	return sum
}
