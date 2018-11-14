/*
	For more informations check:
		https://tour.golang.org/basics/8

*/
package main

import "fmt"

func main() {
	// var statement declares a variable
	var x bool //declare boolean, default value false
	var y int  // declare integer, default value 0

	// you can declare multiple variables at once
	var name, gender string = "Zoran", "Male"

	/*
		The := syntax is shorthand for declaring and initializing a variable
		Note: Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	*/
	years := 29

	//Print result:
	fmt.Println(x, y, years, name, gender)
	// false 0 29 Zoran Male
}
