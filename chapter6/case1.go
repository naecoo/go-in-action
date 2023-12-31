package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 分配一个逻辑处理器
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create goroutine")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("Terminating program")
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer&inner == 0 {
				continue next
			}
			fmt.Printf("%s:%d\n", prefix, outer)
		}
		fmt.Println("Completed", prefix)
	}
}
