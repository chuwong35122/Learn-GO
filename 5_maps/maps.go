package main

import "fmt"

func main() {
	/* Define a map containing the release year of several languages */
	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	/* Loop through each entry and output the name and release year */
	for k, v := range firstReleases {
		fmt.Printf("%s was first released in %d\n", k, v)
	}

	fmt.Println(firstReleases["C"]) //1972
	for key, element := range firstReleases {
		fmt.Println(key)     // C
		fmt.Println(element) //1962
	}

	/* Check if data exists in Map */
	if val, exists := firstReleases["C"]; exists {
		fmt.Println("val ", val)       // val value of "C"
		fmt.Println("exists ", exists) // Will be true of key exists
	}

	if val, exists := firstReleases["foo"]; exists {
		// Note: val will receive value from map
		// exists will be true if key exists

		fmt.Println("val ", val)
		fmt.Println("exists ", exists)
	}
}
