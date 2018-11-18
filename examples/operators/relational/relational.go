package main

import "fmt"

var printIt = fmt.Println

func main() {

	x := 5
	y := 5

	//check if the values of two operands are equal or not;
	printIt(x == y)

	//check if the values of two operands are not equal or not;
	printIt(x != y)

	//check if the value of left operand is greater than the value of right operand;
	printIt(x > y)

	//check if the value of left operand is less than the value of right operand;
	printIt(x < y)

	//check if the value of left operand is greater or equal than the value of right operand;
	printIt(x >= y)

	//check if the value of left operand is less or equal than the value of right operand;
	printIt(x <= y)

	/*
		true
		false
		false
		false
		true
		true
	*/

}
