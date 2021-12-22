package main

import "fmt"

/*
Arrays has fixed size and common data for all element.

*/

func main() {
	var fruits = [4]string{"Banana", "Apple", "Strawberry", "Orange"}
	for i := 0; i < len(fruits); i++ {
		fruit := fruits[i]
		fmt.Println(i, fruit)
	}

	// Array of unassigned numbers
	var unassigned_numbers [5]int
	// Array of numbers
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

}
