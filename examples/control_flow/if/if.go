package main

import "fmt"

var printIt = fmt.Println

func main() {
	number := 0

	//An if statement consists of a boolean expression followed by one or more statements.
	if number < 0 {
		printIt("Number is less than zero")
	} else if number == 0 {
		printIt("Number is equal to zero")
	} else {
		printIt("Number is greater then zero")
	}

	/*
		OUTPUT:
		Number is equal to zero
	*/
}
