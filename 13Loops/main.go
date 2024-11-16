package main

import "fmt"

func main() {
	fmt.Println("Loops in the Go-lang")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "friday", "saturday"}
	fmt.Println(days)

	//
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//
	for i := range days {
		fmt.Println(days[i])
	}

	//
	for ind, day := range days {
		fmt.Printf("The %vth day is %v\n", ind, day)
	}

	//
	rougeVal := 1
	for rougeVal < 10 {

		if rougeVal == 7 {
			goto res
		}

		if rougeVal == 5 {
			rougeVal++
			continue
		}

		if rougeVal == 10 {
			fmt.Println("Found it!")
			rougeVal++
		}

		fmt.Printf("The value is %v\n", rougeVal)
		rougeVal++
	}

res:
	fmt.Println("GoTo statement!")
}
