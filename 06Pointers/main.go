package main

import "fmt"

func main() {

	fmt.Println("Welcome to a class on pointers")

	var ptr *int
	fmt.Println("The default value of pointer is ", ptr)

	myNumber := 23
	var myptr = &myNumber

	fmt.Println("The momory address of the pointer ", myptr)
	fmt.Println("The actual value of the pointer ", *myptr)

	*myptr = *myptr + 2
	fmt.Println("New value is ", myNumber)
}
