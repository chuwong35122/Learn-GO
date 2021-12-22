package main

import "fmt"

// Declaring single variable
var a int

// Declaring multiple variables
/*
Variables are declared followed by its type.
*/
var (
	b bool
	c float32
	d string
)

func main() {
	a = 69
	b, c = true, 420.69
	d = "I'm ready to GO!"

	fmt.Println(a, b, c, d)

	/*
		short variable declaration
		The top code can be re-declared like this...
	*/

	w := 69
	x, y := true, 420.69
	z := "Here we Go again..."

	fmt.Println(w, x, y, z)
}

/*
 The natural value of variable is null

 Types:
 number: 8,16,32,64 bit asssigned using "int" and unsigned using "uint"
	   Floating-point numbers float32, float64
	   Complex number complex64, complex128
 booleans: bool
 string: "Strings" of UTF-8 characters
 errors:
 custom types:
*/
