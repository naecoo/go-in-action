package main

import (
	"fmt"
)

func main() {
	fmt.Println("Chapter 4")

	var arr1 [3]string
	arr2 := [3]string{"1", "2", "3"}

	// 复制一份arr2数据，指向arr1
	// go的数组赋值，会复制一个新的数组（最后传递引用）
	arr1 = arr2
	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	arr1[1] = "Hello World!"
	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)

	arr4 := []int{1, 2, 3, 4, 5}
	arr5 := arr4[1:3:3]             // 第3个参数可以指定容量，默认切片共用同一个底层数组，因此多个切片的修改会互相影响
	arr5 = append(arr5, 10, 11, 14) // 容量不足时，会创建新的底层数组
	fmt.Println("arr4: ", arr4)
	fmt.Println("arr5: ", arr5)

	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7f50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
