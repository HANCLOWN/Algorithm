package main

import (
	"fmt"
	"math"
)

func main() {
	//arr := []int{1, 3, 3, 1, 1, 3, 3, 1}
	//s := 0
	//count := 0
	//wildIdCount := 0
	//realIcons := 0
	//wildId := 1
	//for k, v := range arr {
	//	if k == 0 || s == wildId {
	//		s = v
	//	}
	//	if s == v || v == wildId {
	//		count++
	//		if s == wildId {
	//			wildIdCount++
	//		}
	//		realIcons |= 1 << v
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println("s", s, "count", count, "wildIdCount", wildIdCount, "realIcons", realIcons)
	a := uint32(300)
	s := math.Max(float64(a), 100)
	fmt.Println(s)
}
