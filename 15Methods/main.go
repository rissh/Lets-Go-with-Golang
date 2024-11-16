package main

import "fmt"

func main() {
	fmt.Println("Methods in golnag")
	// No inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 23}
	fmt.Println(hitesh)

	//
	fmt.Printf("hitesh details are: %+v\n", hitesh)

	//
	fmt.Printf("Name is %v and email is %v.\n", hitesh.name, hitesh.Email)

	//Method call
	hitesh.GetStatus()
	hitesh.newEmail()

	fmt.Printf("Name is %v and email is %v.\n", hitesh.name, hitesh.Email)
}

type User struct {
	name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user is male?", u.Status)
}

func (u User) newEmail() {
	u.Email = "test@go.dev"
	fmt.Println("The new email is: ", u.Email)
}
