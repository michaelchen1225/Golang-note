package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	" ایند مالس",
	"Привет, мир",
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(helloList))
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}