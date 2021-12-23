package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		Slices are dynamic arrays.
		Length of it is the current number of elements.
		Capacity is the maximum number of elements.
	*/

	letters := [9]string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", //need comma
	}

	first_four := letters[:4]
	fmt.Println(first_four)
	last_four := make([]string, 4) // Create slice of string with length 4
	last_four = letters[5:]
	fmt.Println(last_four)

	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma
	}
	allLangs := languages[:]                     // copy of the array
	fmt.Println(reflect.TypeOf(allLangs).Kind()) // slice

	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:4:4]         // length 4 capacity 4
	frameworks = append(frameworks, "Meteor") // Function put item at the end of the slide, it is not possible with array.

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js frameworks: %v\n", jsFrameworks)
}
