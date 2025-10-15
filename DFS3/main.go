package main

import "fmt"

// 定义连通块结构体
type Cluster struct {
	Value     int   // 连通块代表的数字
	Size      int   // 连通块大小
	Positions []int // 位置列表 [行,列]
}

func main() {
	board := [][]int{
		{3, 3, 7, 7, 7},
		{9, 1, 3, 3, 5},
		{8, 1, 5, 6, 4},
		{4, 7, 5, 5, 5},
		{1, 7, 1, 9, 6},
		{3, 3, 1, 1, 2},
	}
	N := 5 // 最小连通阈值

	clusters := findClusters(board, N)
	allPositions := make(map[int]bool) // 存储所有消除位置（去重）

	// 输出所有连通块并收集消除位置
	for _, cluster := range clusters {
		fmt.Printf("链接id %d 个数 %d 位置 %v\n",
			cluster.Value, cluster.Size, cluster.Positions)

		for _, pos := range cluster.Positions {
			allPositions[pos] = true
		}
	}

	// 转换并输出所有消除位置
	var positionsList []int
	for pos := range allPositions {
		positionsList = append(positionsList, pos)
	}
	fmt.Println("所有消除的位置：", positionsList)
}

func findClusters(board [][]int, N int) []Cluster {
	rows := len(board)
	cols := len(board[0])
	var clusters []Cluster

	// 获取所有非1的数字种类
	uniqueNumbers := make(map[int]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] != 1 {
				uniqueNumbers[board[i][j]] = true
			}
		}
	}

	// 针对每种数字进行连通块搜索
	for num := range uniqueNumbers {
		visited := make([][]bool, rows)
		for i := range visited {
			visited[i] = make([]bool, cols)
		}

		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				// 只处理当前数字或赖子（1）且未访问的位置
				if (board[i][j] == num || board[i][j] == 1) && !visited[i][j] {
					//var positions [][2]int
					//queue := [][2]int{{i, j}}
					var positions []int
					queue := []int{i*10 + j}
					visited[i][j] = true

					// BFS搜索连通块
					for len(queue) > 0 {
						pos := queue[0]
						queue = queue[1:]
						//x, y := pos[0], pos[1]

						x, y := pos/10, pos%10
						positions = append(positions, pos)

						// 检查四个方向：上、下、左、右
						directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
						for _, d := range directions {
							nx, ny := x+d[0], y+d[1]
							// 检查边界
							if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
								// 只处理当前数字或赖子（1）
								if (board[nx][ny] == num || board[nx][ny] == 1) && !visited[nx][ny] {
									visited[nx][ny] = true
									//queue = append(queue, [2]int{nx, ny})
									queue = append(queue, nx*10+ny)
								}
							}
						}
					}

					// 记录满足条件的连通块
					if len(positions) >= N+1 {
						clusters = append(clusters, Cluster{
							Value:     num,
							Size:      len(positions),
							Positions: positions,
						})
					}
				}
			}
		}
	}
	return clusters
}
