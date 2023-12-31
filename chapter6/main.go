package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// cpu核心数
	fmt.Printf("CPU cores: %d\n", runtime.NumCPU())

	// 分配一个逻辑处理器
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutine")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println()
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("terminating program")
}
