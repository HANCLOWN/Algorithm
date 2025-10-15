package class01

// 选择排序
func SelectSort(arr []int) {
	//0~n-1
	//1~n-1
	//2~n-1
	//3~n-1
	//第一个数的变化由第一个循环控制也就是说几到n-1
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// 冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 插入排序
// 从第一位往前看，如果前面大于则交换
// 从第二位往前看，如果前面大于则交换
// 从第三位往前看，如果前面大于则交换
// 这里之所以可以这么写j := i; j > 0 && arr[j-1] > arr[j]; j--
// 是因为满足条件就换，遇到不满足的停止 因为不满则的之前的肯定是有序的
// 这也就是左神为什么说插入排序可能优秀于冒泡，因为插入可能终止但是冒泡必须全部执行
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}
