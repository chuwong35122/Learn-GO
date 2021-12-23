package main

import "fmt"

func main() {
	a := make([]int, 200+1)
	for i := range a {
		a[i] = i
	}

	addition := a[100] + a[200]
	fmt.Println(a)

	fmt.Println("Addition", addition)

	for i := range a {
		a[i] = a[i] + 100 + 200
	}
	fmt.Println(a)
}
