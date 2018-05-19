package main

import (
	"runtime"
	"fmt"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("memAlloc:", mem.Alloc)
	fmt.Println("memTotalAlloc", mem.TotalAlloc)
	fmt.Println("memHeapAlloc", mem.HeapAlloc)
	fmt.Println("memNumGc", mem.NumGC)
	fmt.Println("....")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}

	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation Failed")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
	//GODEBUG=gctrace=1 go run main.go
}
