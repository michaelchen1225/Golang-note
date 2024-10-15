package main

import "fmt"

var x = "pkg" // package scope

func main() {
	fmt.Println("Main Scope:", x) // read from package scope

	x := "main" // redeclare x in main scope

	fmt.Println("Main Scope:", x) // read from main scope

	if true {
		fmt.Println("If Scope:", x) // read from main scope
		printX()
	}

}

func printX() {
	fmt.Println("Function Scope:", x) // read from package scope
}
