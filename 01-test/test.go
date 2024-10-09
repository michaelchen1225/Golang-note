package main

import "fmt"

// var x *int

func main() {
	a := 5
	// x = 1
	place := &a
	fmt.Println("a's place = ", place)
	// fmt.Println("x = ", x)
	// fmt.Println("*x = ", *x)
}
