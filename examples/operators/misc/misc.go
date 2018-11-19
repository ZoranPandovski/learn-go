/*
	For more info check: https://golangbot.com/pointers/
*/
package main

import "fmt"

func main() {

	// & returns the address of a given variable
	x := 50
	// print x address
	fmt.Println(&x)

	// * is a pointer to a particular variable.
	//The zero value of a pointer is nil
	var y *int
	y = &x
	//print y value
	fmt.Println(*y)

	//deferencing a pointer value
	*y++
	fmt.Println(x)
	fmt.Println(*y)

	/*
		OUTPUT:
		0xc000088010
		50
		51
		51
	*/
}
