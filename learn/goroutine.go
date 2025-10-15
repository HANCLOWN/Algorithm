package learn

import (
	"fmt"
	"time"
)

func GoroutineGo() {
	//定义个输出奇数的函数
	printOdd := func() {
		for i := 1; i <= 10; i = i + 2 {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}

	}
	//	定义个输出偶数的函数
	printEven := func() {
		for i := 2; i <= 10; i = i + 2 {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}

	}
	go printOdd()
	go printEven()
	time.Sleep(1 * time.Second)

}
