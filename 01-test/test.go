package main

import "fmt"

var sum int = 0

func main() {

	for i := range 11 {
		sum += i
	}
	fmt.Println(sum)

}
