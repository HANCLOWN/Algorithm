package main

import (
	"fmt"
)

var rows, cols int
var directions = []struct{ dx, dy int }{
	{0, 1},  // 向右
	{1, 0},  // 向下
	{0, -1}, // 向左
	{-1, 0}, // 向上
}

func dfs(board [][]int, visited [][]bool, x, y int, number int) (count int, positions [6]int) {
	// 检查x, y是否越界
	if x < 0 || x >= cols || y < 0 || y >= len(board[x]) || visited[x][y] || board[x][y] != number {
		return 0, [6]int{0, 0, 0, 0, 0, 0}
	}

	visited[x][y] = true
	count = 1
	positions[x] = 1 << y
	for _, dir := range directions {
		nextX, nextY := x+dir.dx, y+dir.dy
		// 确保下一个位置不会越界
		nextCount, nextPositions := dfs(board, visited, nextX, nextY, number)
		if nextCount != 0 {
			count += nextCount
			for m, v := range nextPositions {
				positions[m] = positions[m] | v
			}
		}
	}
	return count, positions
}

func eliminateNumbersModified(board [][]int, n int) (map[int]int, [6]int) {
	cols = len(board)
	rows = 5
	//辅助标记
	visited := make([][]bool, cols)
	for i := range visited {
		visited[i] = make([]bool, rows)
	}
	res := [6]int{}
	iconCount := make(map[int]int)
	for i := 0; i < cols; i++ {
		for j := 0; j < len(board[i]); j++ {
			if !visited[i][j] && board[i][j] != 0 {
				number := board[i][j]
				count, positions := dfs(board, visited, i, j, number)
				if count >= n {
					iconCount[number] = count
					for m, v := range positions {
						res[m] = res[m] | v
					}
				}
			}
		}
	}
	return iconCount, res
}

func main() {
	board := [][]int{
		// 定义每一列的数字情况
		{1, 1, 7, 7, 7},
		{1, 1, 30, 31, 3},
		{3, 1, 5, 4, 4},
		{4, 4, 5, 5, 5},
		{1, 5, 61, 62, 63},
		{6, 6, 21, 22, 23},
	}

	kindCount, eliminatedCoords := eliminateNumbersModified(board, 3)

	fmt.Println("消除的数字种类数量:", kindCount)
	fmt.Println("应该被消除的坐标:", eliminatedCoords)
}
