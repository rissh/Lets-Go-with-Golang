package main

import "fmt"

func main() {
	fmt.Println("Functions in Go-lang")
	//
	greet()

	//
	res := adder(3, 5)
	fmt.Println("The result is: ", res)

	//
	proRes := proAdder(1, 2, 3, 4, 5)
	fmt.Println("The pro result is: ", proRes)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}

	return total
}

func greet() {
	fmt.Println("Hey man!")
}
