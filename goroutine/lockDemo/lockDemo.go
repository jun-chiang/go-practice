package lockDemo

import (
	"fmt"
	"sync"
	"time"
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
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	fmt.Println("The result of add with lock:", x)
	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	fmt.Println("The result of add without lock:", x)
}
