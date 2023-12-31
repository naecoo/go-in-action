package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func incCount(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func main() {
	wg.Add(2)

	go incCount(1)
	go incCount(2)

	wg.Wait()
	fmt.Println("Final counter value: ", counter)
}
