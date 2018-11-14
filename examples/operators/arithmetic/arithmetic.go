package main

import "fmt"

var printIt = fmt.Println

func main() {
	x := 6
	y := 3

	var sum = x + y // + Adds two operands
	printIt(sum)

	var sub = x - y // - Subtracts second operand from the first
	printIt(sub)

	var pow = x * y // * Multiplies both operands
	printIt(pow)

	var div = x / y // / Divides the numerator by the denominator.
	printIt(div)

	var mo = x % y // % Modulus operator; gives the remainder after an integer division.
	printIt(mo)

	var z int = 5
	z++ //Increment operator. It increases the integer value by one.
	printIt(z)
	z-- // Decrement operator. It decreases the integer value by one.
	printIt(z)

	/*OUTPUT:
	9
	3
	18
	2
	0
	6
	5
	*/
}
