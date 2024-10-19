package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s! Welcome to Go.\n", name)

	var x int = 5
	fmt.Println(x)
	fmt.Printf("%d\n", x)

	var height float32 = 5.11
	fmt.Println("Height:", height)
	fmt.Printf("Height:%f\n", height)

	place := "Mars"
	fmt.Println("I am from", place)

	val := 100

	if val > 100 {
		fmt.Println("Val is greater than 100")
	} else {
		fmt.Println("Val is NOT greater than 100")
	}	

	// loops

	for i := 0; i < 5; i ++ {
		fmt.Println(i)
	}

	i := 10
	for i < 15 {
		fmt.Println(i)
		i++
	}

	// array
	nums := []int{1,2,3,4,5}
	fmt.Println(nums)

	for idx, val := range nums {
		fmt.Println(idx, val)
	}


}