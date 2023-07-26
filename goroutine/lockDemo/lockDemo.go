package lockDemo

import (
	"fmt"
	"sync"
)

// 共享内存实现通信
var (
	x    int32 = 0
	lock sync.Mutex
)

func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
}

func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x++
	}
}

// 测试函数 测试加锁与否的效果
func AddLockDemo() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			addWithLock()
		}()
	}
	wg.Wait()
	fmt.Println("The result of add with lock:", x)
	x = 0
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			addWithoutLock()
		}()
	}
	wg.Wait()
	fmt.Println("The result of add without lock:", x)
}
