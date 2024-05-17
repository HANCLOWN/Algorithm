package main

import (
	"Algorithm/framework/class01"
	"fmt"
)

func main() {
	arr := []int{6, 2, 8, 10, 5, 1, 7, 3, 9, 4}
	class01.BubbleSort(arr)
	fmt.Println(arr)
}
