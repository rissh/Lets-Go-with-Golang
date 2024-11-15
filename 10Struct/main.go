package main

import "fmt"

func main() {
	fmt.Println("Struct in golnag")
	// No inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 23}
	fmt.Println(hitesh)

	//
	fmt.Printf("hitesh details are: %+v\n", hitesh)

	//
	fmt.Printf("Name is %v and email is %v.\n", hitesh.name, hitesh.Email)
}

type User struct {
	name   string
	Email  string
	Status bool
	Age    int
}
