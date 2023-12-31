package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int

func incCount() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}

func main() {

	wg.Add(2)

	go incCount()
	go incCount()

	wg.Wait()
	fmt.Println("Final counter value: ", counter)
}
