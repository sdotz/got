// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

// Add imports.

import "fmt"

// main is the entry point for the application.
func main() {
	var strings [5]string

	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	strings = fruits

	// Assign the populated array to the array of zero values.		

	// Iterate over the first array declared.
	
	for i, str := range strings {
		fmt.Println(i, str, &strings[i])
	}
	
	// Display the string value and address of each element.
}