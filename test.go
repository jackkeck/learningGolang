package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")
	foo()
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			n, _ := fmt.Println(i)
			fmt.Println(n)
		}
	}
}

func foo() {
	fmt.Println("i am so excited right now..")
}
