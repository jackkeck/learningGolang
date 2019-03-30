package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello world!")
	foo()
	foo()

	for i := 0; i < 100; i++ {
		fmt.Println("I am looping " + strconv.Itoa(i) + " times.")
	}
}

func foo() {
	fmt.Println("i am so excited right now..")
}
