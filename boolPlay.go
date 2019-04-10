package main

import "fmt"

var boolTest bool

func main() {
	fmt.Println("boolTest is: ", boolTest)

	boolTest = false

	fmt.Println("boolTest is now: ", boolTest)

	if boolTest == false {
		fmt.Println("We have set boolTest to false and fell into our if")
	}

	a := 200
	b := 180

	fmt.Println("We have set our variables a and b")

	for a != b {
		fmt.Println("a is not equal to b therefore this returned true")
		fmt.Println("a is ", a, " and b is ", b)
		b += 1
	}

	fmt.Println("After our for loop, we have now incremented b to ", b)

}
