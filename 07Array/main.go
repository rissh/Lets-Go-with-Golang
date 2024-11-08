package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Array class")

	var fruitsList [4]string

	fruitsList[0] = "Apple"
	fruitsList[1] = "Banana"
	fruitsList[3] = "Peach"

	fmt.Println("The Fruit List is ", fruitsList)
	fmt.Println("The lenght od list is ", len(fruitsList))
	fmt.Println("The Capacity of the list is ", cap(fruitsList))

	var vegList = [3]string{"Onion", "Tomato", "Cabbage"}

	fmt.Println("The vegetavles list is ", vegList)
	fmt.Println("The lenght os vegetables list is ", len(vegList))
	fmt.Println("The Capacity of the list is ", cap(vegList))
}
