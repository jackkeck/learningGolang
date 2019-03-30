package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 30
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 42
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
