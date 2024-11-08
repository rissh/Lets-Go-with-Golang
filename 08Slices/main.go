package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices class")

	var fruitList = []string{"Apple", "Peach"}
	fmt.Printf("Type of the fruits list is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// Make() function

	highscores := make([]int, 5)

	highscores[0] = 234
	highscores[1] = 567
	highscores[2] = 890
	highscores[3] = 987
	highscores[4] = 654

	highscores = append(highscores, 435, 648, 957)
	fmt.Println(highscores)

	// Inbuild methods
	sort.Ints(highscores)
	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))

	// How to remove a value from slices based on index
	var courses = []string{"react", "next", "python", "java", "c++", "ruby", "swift"}
	fmt.Println(courses)

	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
