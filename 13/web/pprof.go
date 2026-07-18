package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	profiledemo "example.com/hello/13/13.8"
)

var (
	checksum   atomic.Uint64
	retainedMu sync.Mutex
	retained   [][]byte
)

func main() {
	http.HandleFunc("/work", workHandler)
	http.HandleFunc("/alloc", allocHandler)
	log.Println("服务器：http://localhost:6060")
	log.Println("pprof：http://localhost:6060/debug/pprof/")
	log.Println("CPU 工作：http://localhost:6060/work?seconds=10")
	log.Println("内存分配：http://localhost:6060/alloc?mb=64")

	if err := http.ListenAndServe("localhost:6060", nil); err != nil {
		log.Fatal(err)
	}
}

func workHandler(w http.ResponseWriter, r *http.Request) {
	seconds := 10
	if value := r.URL.Query().Get("seconds"); value != "" {
		number, err := strconv.Atoi(value)
		if err != nil || number < 1 || number > 60 {
			http.Error(w, "seconds must between 1 and 60", http.StatusBadRequest)
			return
		}
		seconds = number
	}
	deadline := time.Now().Add(time.Duration(seconds) * time.Second)

	var res uint64

	for time.Now().Before(deadline) {
		res ^= profiledemo.CPUWork(200_000)
	}

	checksum.Store(res)
	fmt.Fprintf(w, "CPU 工作完成:seconds=%d checksum=%d\n", seconds, res)
}

func allocHandler(writer http.ResponseWriter, request *http.Request) {
	megabytes := 64
	if value := request.URL.Query().Get("mb"); value != "" {
		number, err := strconv.Atoi(value)
		if err != nil || number < 1 || number > 512 {
			http.Error(writer, "mb 必须是 1 到 512 的整数", http.StatusBadRequest)
			return
		}
		megabytes = number
	}

	blocks := profiledemo.AllocateMB(megabytes)

	retainedMu.Lock()
	retained = blocks
	retainedMu.Unlock()

	fmt.Fprintf(writer, "已经分配并保留 %d MiB\n", megabytes)
}
