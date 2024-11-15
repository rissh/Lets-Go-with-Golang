package main

import "fmt"

func main() {
	fmt.Println("Mas in Golang!")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of the languages: ", languages)
	fmt.Println("JS stands for: ", languages["JS"])

	//
	delete(languages, "RB")
	fmt.Println("List of the languages: ", languages)

	//Loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
