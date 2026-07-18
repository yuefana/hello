package profiledemo

// CPUWork 执行一段纯 CPU 计算，用于产生可观察的 CPU profile。
func CPUWork(rounds int) uint64 {
	x := uint64(0x9e3779b97f4a7c15)

	for i := 0; i < rounds; i++ {
		x ^= uint64(i) + 0x9e3779b97f4a7c15
		x = (x << 7) | (x >> 57)
		x *= 0xbf58476d1ce4e5b9
	}

	return x
}

// AllocateMB 分配并实际访问指定大小的内存。
func AllocateMB(megabytes int) [][]byte {
	if megabytes <= 0 {
		return nil
	}

	blocks := make([][]byte, megabytes)

	for i := range blocks {
		block := make([]byte, 1<<20) // 1 MiB

		// 每隔一个内存页写一次，避免只保留虚拟地址而没有实际使用。
		for offset := 0; offset < len(block); offset += 4096 {
			block[offset] = byte(i)
		}

		blocks[i] = block
	}

	return blocks
}
