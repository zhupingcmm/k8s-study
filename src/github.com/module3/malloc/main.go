package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int

	for {
		_ = make([]byte, 1024*1024*10) // 分配1MB的内存
		counter++
		fmt.Printf("Allocated %d MB of memory. Total allocations: %d\n", counter*10, counter*10)
		time.Sleep(time.Second) // 每秒分配一次
	}
}
