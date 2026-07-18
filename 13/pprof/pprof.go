package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	profiledemo "example.com/hello/13/13.8"
)

var (
	cpuProfile = flag.String("cpuprofile", "", "将cpu profile 写入指定文件")
	memProfile = flag.String("memprofile", "", "将内存profile写入文件")
	memoryMB   = flag.Int("memmb", 64, "分配并保留的内存大小,单位MB")
	seconds    = flag.Int("seconds", 3, "Cpu工作持续秒数")
)

func main() {
	fmt.Printf("解析前 os.Args：%q\n", os.Args)

	flag.Parse()

	fmt.Printf("cpuProfile = %q\n", *cpuProfile)
	fmt.Printf("memProfile = %q\n", *memProfile)
	fmt.Printf("seconds = %d\n", *seconds)
	fmt.Printf("memoryMB = %d\n", *memoryMB)
	fmt.Printf("未解析的参数 = %q\n", flag.Args())
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if *seconds <= 0 {
		return errors.New("second must > 0")
	}
	if *memoryMB < 0 {
		return errors.New("memorMb is neg")
	}
	if *cpuProfile != "" {
		file, err := os.Create(*cpuProfile)
		if err != nil {
			return fmt.Errorf("创建 CPU profile失败%w", err)
		}
		defer file.Close()
		if err := pprof.StartCPUProfile(file); err != nil {
			return fmt.Errorf("启动cpu profile 失败%w", err)
		}
		defer pprof.StopCPUProfile()
	}
	deadline := time.Now().Add(time.Duration(*seconds) * time.Second)

	var checksum uint64

	for time.Now().Before(deadline) {
		checksum ^= profiledemo.CPUWork(200_000)
	}

	blocks := profiledemo.AllocateMB(*memoryMB)

	if *memProfile != "" {
		file, err := os.Create(*memProfile)
		if err != nil {
			return fmt.Errorf("创建内存 profile 失败：%w", err)
		}
		defer file.Close()
		runtime.GC()

		profile := pprof.Lookup("allocs")
		if profile == nil {
			return errors.New("找不到 allcos profile")
		}
		if err := profile.WriteTo(file, 0); err != nil {
			return errors.New("写入内存失败")
		}
	}
	runtime.KeepAlive(blocks)
	fmt.Printf("运行完成：checksum=%d，保留内存=%d MiB\n", checksum, len(blocks))

	return nil
}
