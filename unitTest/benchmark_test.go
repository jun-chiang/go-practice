package unittest

import "testing"

func BenchmarkSelect(b *testing.B) {
	// 初始化服务器列表
	InitServerIndex()
	// 上面的操作不属于测试范围，所以重置计时器
	b.ResetTimer()
	// 这个N是由框架自动设定，以便在测试运行足够长的时间后产生有意义的结果
	for i := 0; i < b.N; i++ {
		Select()
	}
}

func BenchmarkSelectParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	// 并发执行基准测试
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Select()
		}
	})
}
