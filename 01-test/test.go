package main

import "fmt"

func main() {
	str1 := "你好"
	// myrune := []rune(str1)

	for i, v := range str1 {
		fmt.Printf("index: %d, value: %c, type: %T\n", i, v, v)
	}
}
