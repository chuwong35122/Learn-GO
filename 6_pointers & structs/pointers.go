package main

import "fmt"

/*
Pointers arithmetics are not allowed in Go.
It is stored in memory address using type *T (a pointer of T).
Default value is nil.
*/

func main() {
	var address *int
	number := 420

	address = &number // address stores memory address
	value := *address // get the value from pointer

	fmt.Println("address ", address)
	fmt.Println("value ", value)
}
