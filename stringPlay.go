package main

import "fmt"

var s string

func main() {

	s = "this is a string."

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

}
