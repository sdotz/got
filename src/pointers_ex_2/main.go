// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.

import "fmt"
// Declare a type named user.

type user struct {
	name string
	age int 
}

// Create a function that changes the value of one of the user fields.
func funcName( usr *user, age int) {
	usr.age = age
}

// main is the entry point for the application.
func main() {
 	me := user {name: "stephen", age:26}

	fmt.Println(me.name, me.age)

	funcName(&me, 27)

	// Display the value of the variable.
	fmt.Println(me.name, me.age)
}