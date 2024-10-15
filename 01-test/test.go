package main

import (
	"errors"
	"fmt"
	"log"
)

var x int

func inputAndCheck(input int) error {

	if input < 7 || input > 20 {
		return errors.New("輸入錯誤，請輸入 7 ~ 20 之間的整數")
	} else {
		return nil
	}
}

func main() {
	fmt.Print("請輸入一個介於 7 ~ 20 的整數：")
	_, error := fmt.Scanf("%d", &x)

	if error != nil {
		fmt.Println("輸入錯誤，請輸入整數")
		log.Fatal(error)
	} else if error := inputAndCheck(x); error != nil {
		fmt.Println(error)
	} else if x%2 == 0 {
		fmt.Println("偶數")
	} else {
		fmt.Println("奇數")
	}
}
