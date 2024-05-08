package model

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 数组拷贝器
func CopyArray(a []int) []int {
	arr := make([]int, len(a))
	for i, _ := range arr {
		arr[i] = a[i]
	}
	return arr
}

// 随机生成长度数值范围的数组
func LenRandomValueRandom(maxLen int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arrLen := int(rand.Float64() * float64(maxLen))
	arr := make([]int, arrLen)
	for i, _ := range arr {
		arr[i] = int(rand.Float64() * float64(maxValue))
	}
	return arr
}

// 随机生成长度数值范围的数组相邻不等
func LenRandomValueRandomNoDENG(maxLen int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arrLen := int(rand.Float64() * float64(maxLen))
	arr := make([]int, arrLen)
	for i, _ := range arr {
		arr[i] = int(rand.Float64() * float64(maxValue))
		for i > 0 && arr[i] == arr[i-1] {
			arr[i] = int(rand.Float64() * float64(maxValue))
		}
	}
	return arr
}

// 内存分配查看器
func PrintMemUsage() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	// 打印以字节为单位的内存使用情况
	fmt.Printf("Alloc = %v MiB", bToMb(mem.Alloc))             // 已分配并仍在使用的字节
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(mem.TotalAlloc)) // 自程序启动以来分配的总字节数
	fmt.Printf("\tSys = %v MiB", bToMb(mem.Sys))               // 从操作系统获取的内存
	fmt.Printf("\tNumGC = %v\n", mem.NumGC)                    // 垃圾收集的次数
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// 用法
func main() {
	PrintMemUsage() // 打印开始运行时的内存使用情况

	// 这里写你的程序...

	PrintMemUsage() // 打印程序运行后的内存使用情况
}

// generateRandomInt2DArray 生成一个 n*m 的二维数组，并用指定范围 [min, max) 的随机整数填充
func GenerateRandomInt2DArray(n, m, min, max int) [][]int {
	rand.Seed(time.Now().UnixNano())
	array := make([][]int, n)
	for i := 0; i < n; i++ {
		array[i] = make([]int, m)
		for j := 0; j < m; j++ {
			array[i][j] = rand.Intn(max-min) + min
		}
	}
	return array
}

// generateRandomFloat642DArray 生成一个 n*m 的二维数组，并用指定范围 [min, max) 的随机浮点数填充
func GenerateRandomFloat642DArray(n, m int, min, max float64) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	array := make([][]float64, n)
	for i := 0; i < n; i++ {
		array[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			array[i][j] = rand.Float64()*(max-min) + min
		}
	}
	return array
}

// generateRandomUint322DArray 生成一个 n*m 的二维数组，并用指定范围 [min, max) 的随机浮点数填充
func GenerateRandomUint322DArray(n, m int, min, max int) [][]uint32 {
	rand.Seed(time.Now().UnixNano())
	array := make([][]uint32, n)
	for i := 0; i < n; i++ {
		array[i] = make([]uint32, m)
		for j := 0; j < m; j++ {
			array[i][j] = uint32(rand.Intn(max-min) + min)
		}
	}
	return array
}
