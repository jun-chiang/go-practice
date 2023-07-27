package unittest

import "math/rand"

// 模拟服务器列表
var ServerIndex [10]int

// 初始化服务器列表
func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

// 随机选择服务器
func Select() int {
	return ServerIndex[rand.Intn(10)]
}
