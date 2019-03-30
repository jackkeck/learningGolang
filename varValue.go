package main

import "fmt"

var y = 22
var z int

func main() {

	//short declaration
	x := 2

	//var declaration can live outside of function.
	fmt.Println(x + y + z)
	z += 1
	fmt.Println(x + y + z)
}
