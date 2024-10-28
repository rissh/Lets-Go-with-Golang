package main

import "fmt"

// Public const variable
const LoginToken string = "ghabbhhjd"

func main() {
	fmt.Println("Variables")

	var username string = "Rishi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isloggin bool = true
	fmt.Println(isloggin)
	fmt.Printf("Variable is of type: %T \n", isloggin)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallFloat float32 = 255.343434343434
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.34343434343434
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	// Default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// Implicit type

	var web = "Little Monster"
	fmt.Println(web)
	fmt.Printf("Variable is of type: %T \n", web)

	// No var style

	numberOfUser := 23022001
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	//

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
