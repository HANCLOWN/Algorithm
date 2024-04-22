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

func dfs(board [][]int, visited [][]bool, x, y int, number int) (count int, positions []struct{ x, y int }) {
	if x < 0 || y < 0 || x >= rows || y >= cols || visited[x][y] || board[x][y] != number {
		return 0, nil
	}
	visited[x][y] = true
	count = 1
	positions = append(positions, struct{ x, y int }{x, y})

	for _, dir := range directions {
		nextCount, nextPositions := dfs(board, visited, x+dir.dx, y+dir.dy, number)
		count += nextCount
		positions = append(positions, nextPositions...)
	}
	return count, positions
}

func eliminateNumbers(board [][]int, n int) map[int]int {
	rows, cols = len(board), len(board[0])
	//用于标记的
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	result := make(map[int]int) // 记录消除的数字和对应的消除次数

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] && board[i][j] != 0 { //没被标记并且有效
				number := board[i][j] // 记录当前数字
				count, positions := dfs(board, visited, i, j, number)
				if count >= n {
					for _, pos := range positions {
						board[pos.x][pos.y] = 0 // 消除数字
					}
					result[number] += count // 使用记录的数字而不是后来可能变为0的值
				}
			}
		}
	}

	return result
}

func main() {
	board := [][]int{
		{1, 1, 3, 4, 1, 6},
		{1, 1, 1, 4, 5, 6},
		{7, 30, 5, 5, 61, 21},
		{7, 31, 4, 5, 62, 22},
		{7, 3, 4, 5, 63, 23},
	}

	result := eliminateNumbers(board, 4)

	fmt.Println("消除的数字及其个数:", result)
}
