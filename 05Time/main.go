package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2001 Monday"))

	//
	createdDate := time.Dae(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
