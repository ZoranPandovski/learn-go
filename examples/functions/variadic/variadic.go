package main

import "fmt"

/*
By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters.
For more info: https://gobyexample.com/variadic-functions
*/

//add takes an arbitrary number of ints as arguments.
func add(args ...int) int {

	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(5, 6, 1))
}
