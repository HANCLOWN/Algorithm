package model

import (
	"math/rand"
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
