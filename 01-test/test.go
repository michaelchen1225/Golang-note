package main

import "fmt"

func main() {

	for i, c := range "hello" {
		fmt.Printf("%d: %c\n", i, c)
	}

}
