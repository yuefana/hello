package profiledemo

import "testing"

// 包级变量用于保存结果，防止编译器认为计算结果没有被使用。
var (
	cpuBenchmarkResult    uint64
	memoryBenchmarkResult [][]byte
)

func BenchmarkCPUWork(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		cpuBenchmarkResult = CPUWork(100_000)
	}
}

func BenchmarkAllocate(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		memoryBenchmarkResult = AllocateMB(1)
	}
}
