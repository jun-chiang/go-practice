package channelDemo

import "fmt"

func CalSquare() {
	src := make(chan int)
	dest := make(chan int, 3)

	// 子协程A生产数字
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	// B协程消费数字，计算平方
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()
	// 主线程里面获取计算好的平方数字
	for i := range dest {
		fmt.Println(i)
	}
}
