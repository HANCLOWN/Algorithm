package main

import (
	"Algorithm/model"
	"fmt"
	"math/rand"
)

func main() {
	////前缀和数组29.00
	//elements := []int{1, 2, 3, 4, 5}
	//a := 3
	//c := 0
	//
	//// 调用ProcessElements函数，传递elements数组和一个匿名函数作为参数。
	//// 这个匿名函数将数组中的每个元素加倍。
	//ProcessElements(elements, func(i, j int, iconId int) {
	//	fmt.Println(((1 << i) & a) > 0)
	//	if ((1 << i) & a) > 0 {
	//		return
	//	}
	//	if i > 2 {
	//		c++
	//	}
	//})
	//fmt.Println(c)
	//fmt.Println("==============")
	//realIconId := uint64(0)
	//realIconId = (1 << 5) | realIconId
	//awardIconId := bits.TrailingZeros64(realIconId)
	//fmt.Println(realIconId, awardIconId)
	//fmt.Println("==============")
	//zhou1 := []int{1, 4, 7, 8, 5, 2, 3, 6, 9}
	//zhou2 := []int{7, 8, 9, 6, 5, 4, 1, 2, 3}
	//zhou3 := []int{6, 5, 4, 7, 8, 1, 2, 3, 9}
	//zhou4 := []int{5, 2, 6, 8, 4, 3, 9, 1, 7}
	//zhou5 := []int{4, 8, 6, 2, 1, 7, 9, 3, 5}
	//SizeLine := []int{3, 3, 3, 3, 3}
	//replaceIcon := []int{10, 10, 10, 11, 11, 11, 12, 12, 12, 13, 13, 13, 11, 11, 11}
	////bonusId :=11
	//resul := []int{1, 2, 3, 2, 1}
	//arrs := [][]int{}
	//arrs = append(arrs, zhou1)
	//arrs = append(arrs, zhou2)
	//arrs = append(arrs, zhou3)
	//arrs = append(arrs, zhou4)
	//arrs = append(arrs, zhou5)
	//for i := 0; i < len(SizeLine); i++ {
	//	//这里没+1是因为配置不会从0开始
	//	res := AxisGetter(resul, arrs, i, SizeLine[i])
	//	fmt.Println(res)
	//}
	//fmt.Println("==============////////////////")
	////num := CountIcon(resul, SizeLine, arrs, 4000, 0)
	////本段处理 icons[0][index] == bonusId index = index - index%SizeLine[i]
	////下段处理index = ((index + index%SizeLine[i])+1)/len(15)
	//fmt.Println(ReplaceColums(resul, SizeLine, arrs, 13, 0x0e, replaceIcon, true, false))
	////fmt.Println(num)
	//fmt.Println(1 % 3000)
	////计算全线排列的函数
	//Axis := []int{3, 3, 3, 3, 3}
	//var count = 1
	//for _, v := range Axis {
	//	count *= v
	//}
	//columSize := len(Axis)
	//for k := range count {
	//	item := make([]int, columSize)
	//	//pre := 1
	//	pre := k
	//	for i, v := range Axis {
	//		//pre *= v
	//		//item[i] = (k/(count/pre))%v + 1
	//		item[i] = pre%v + 1
	//		pre /= v
	//	}
	//	fmt.Println("==============", k)
	//	fmt.Println("==============", item)
	//	if k > 10 {
	//		break
	//	}
	//}
	//fmt.Println("==============")
	//index := 100
	//for i := range index {
	//	i = i - i%3
	//	fmt.Println(i)
	//}
	//OneMinIndexMain()
	//rand.New(rand.NewSource(time.Now().UnixNano() + 1))
	//T := []int{30000, 2200, 1200, 475, 60}
	//
	//testTimes := 30
	//spin := 100
	//a = 0
	//for l := range spin {
	//	count = 0
	//	for i := range testTimes {
	//		val := rand.Float32() * float32(30000+2200+1200+475+60)
	//		sum := float32(0)
	//		for k, weight := range T {
	//			sum += float32(weight)
	//			if val <= sum {
	//				if k > 0 {
	//					count++
	//				}
	//				c = i
	//				//fmt.Println("第", i, "次", "结果", k)
	//				break
	//			}
	//		}
	//	}
	//	c = l
	//	if (float64(count) / float64(testTimes)) >= 0.117 {
	//		a++
	//	}
	//	//fmt.Println("第", l, "次", "结果", float64(count)/float64(testTimes)*30)
	//
	//}
	//fmt.Println("结果", a)
	//
	//ans := oneMinIndex([]int{10, 18, 10, 4, 15, 14, 7})
	//fmt.Println("==============", ans)
	//y := Y{}
	//y.method() // 调用 Y 中嵌套的 X 类型的 method，X 仍然是接收者
	model.PrintMemUsage()
}

type X struct{}

func (x X) method() {
	fmt.Println("X's method")
}

type Y struct {
	X
}

func AxisGetter(iconInfo []int, icons [][]int, i, j int) int {
	return icons[i][(iconInfo[i])+j-1%len(icons[i])]
}

func IterateIcons(resul []int, SizeLine []int, icons [][]int, callback func(i, j int, iconId int)) {
	for i := range resul { //列 index [0123]
		for j := range SizeLine[i] { //行 ,偏移量 value 3 3 3 3
			//这里+1是因为rang从0开始的
			callback(i, int(j), AxisGetter(resul, icons, i, j+1))
		}
	}
}

// 统计某个图标的个数
func CountIcon(resul []int, SizeLine []int, icons [][]int, targetIconId int, omits uint32) uint32 {
	count := uint32(0)
	IterateIcons(resul, SizeLine, icons, func(i, j int, iconId int) {
		if ((1 << i) & omits) > 0 {
			return
		}
		if iconId == targetIconId {
			count++
		}
	})
	return count
}

func ReplaceColums(resul []int, SizeLine []int, icons [][]int, index int, colums uint32, replaceIcon []int, turn bool, b bool) any {
	ret := make([][]int, len(SizeLine))
	for i, v := range SizeLine {
		ret[i] = make([]int, v)
	}
	IterateIcons(resul, SizeLine, icons, func(i, j int, iconId int) {
		iss := index
		if turn {
			if b {
				iss = iss - iss%3
				fmt.Println(iss)
			} else {
				iss = (iss + (3 - iss%3) + 15) % 15
				fmt.Println(iss)
			}
		}
		if ((1 << i) & colums) > 0 {
			ret[i][j] = replaceIcon[(iss+j)%len(replaceIcon)]
		} else {
			ret[i][j] = iconId
		}
	})
	return ret
}

func ProcessElements(elements []int, callback func(i, j int, iconId int)) {
	for _, element := range elements {
		// 对每个元素调用回调函数，并打印结果。
		callback(element, 1000, 5)
	}
}

func RandomValues[T uint32 | float32](values []T) uint32 {
	total := T(0)
	for _, v := range values {
		total += v
	}
	val := rand.Float32() * float32(total)
	sum := float32(0)
	for i, weight := range values {
		sum += float32(weight)
		if val <= sum {
			return uint32(i)
		}
	}
	return 0
}

// 二分实现(一个有序数组中有没有指定数字) [1,2,2,3,5,6,8,9] 4-false 5-true
func find(arr []int, num int) bool {
	if arr == nil || len(arr) < 1 {
		return false
	}
	R := 0
	L := len(arr) - 1
	for L <= R {
		mid := (R + L) / 2
		if arr[mid] == num {
			return true
		} else if arr[mid] < num {
			L = mid + 1
		} else if arr[mid] > num {
			R = mid - 1
		}
	}
	return false
}

// 在有序数组中找到>=num的最左位置 [1,2,2,3,5,6,8,9]  100->-1 5->-4 2->1
func MostLestNoLessNumIndex(arr []int, num int) int {
	ans := -1
	if arr == nil || len(arr) < 1 {
		return -1
	}
	R := 0
	L := len(arr) - 1
	for L <= R {
		mid := (R + L) / 2
		if arr[mid] >= num {
			ans = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return ans
}

// 在有序数组中找到<=num的最右面位置 [10,20,20,30,50,60,80,90]  4->-1 50->-4 20->0
func MostRightNoMoreNumIndex(arr []int, num int) int {
	ans := -1
	if arr == nil || len(arr) < 1 {
		return -1
	}
	R := 0
	L := len(arr) - 1
	for L <= R {
		mid := (R + L) / 2
		if arr[mid] <= num {
			ans = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return ans
}

func test(t int, arr []int, num int) any {
	switch t {
	case 1: //find 测试
		if arr == nil || len(arr) < 1 {
			return false
		}
		for _, v := range arr {
			if v == num {
				return true
			}
		}
		return false
	case 2: //MostLestNoLessNumIndex 测试
		if arr == nil || len(arr) < 1 {
			return -1
		}
		for i, v := range arr {
			if v >= num {
				return i
			}
		}
	case 3: //MostRightNoMoreNumIndex 测试
		if arr == nil || len(arr) < 1 {
			return -1
		}
		ans := -1
		for i, v := range arr {
			if v <= num {
				ans = i
			}
			return ans
		}
	}
	return "type not find"
}

// 无序数组局部最小 两个相邻元素不等
func oneMinIndex(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	n := len(arr)
	if len(arr) == 1 {
		return 0

	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[n-1] < arr[n-2] {
		return n - 1
	}
	L := 0
	R := n - 1
	for L < R-1 { //for L <= R {注释部分如果只剩下两个值时候mid-1会越界，改成R-1的话，保证满足条件的情况下至少有L R-1 R三个值，如果跳出循环再进行最后两个值的比较【3,2,3,2,3】
		mid := (L + R) / 2
		if arr[mid] < arr[mid-1] && arr[mid] < arr[mid+1] {
			return mid
		} else {
			if arr[mid] > arr[mid-1] {
				R = mid - 1
			} else {
				L = mid + 1
			}
		}
	}
	if arr[L] < arr[R] {
		return L
	} else {
		return R
	}
}

func checkOneMinIndex(arr []int, index int) bool {
	if len(arr) == 0 {
		return index == -1
	}
	L := index - 1
	R := index + 1
	if (L >= 0 && arr[L] < arr[index]) || (R <= len(arr)-1 && arr[R] < arr[index]) {
		return false
	}
	return true
}

func OneMinIndexMain() {
	maxLen := 10
	maxValue := 20
	testTimes := 100000
	for i := range testTimes {
		arr := model.LenRandomValueRandomNoDENG(maxLen, maxValue)
		ans := oneMinIndex(arr)
		if !checkOneMinIndex(arr, ans) {
			fmt.Println("第", i, "次", "出错啦", arr)
			fmt.Println("第", i, "次", "出错啦", ans)
			break

		} else {
			fmt.Println("第", i, "次", "通过了", arr)
			fmt.Println("第", i, "次", "通过了", ans)

		}
	}
}
