package main

import (
	"github.com/jun-chiang/go-practice/goroutine/lockDemo"
)

func main() {
	// 通道测试的方法
	//channelDemo.CalSquare()
	// 加锁测试的方法
	lockDemo.AddLockDemo()
}
