package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	//Called Logical AND operator. If both the operands are false, then the condition becomes false.
	if a && b {
		fmt.Printf("Condition is true\n")
	}

	//Called Logical OR Operator. If any of the two operands is true, then the condition becomes true.
	if a || b {
		fmt.Printf("Condition is true\n")
	}

	a = false
	b = true
	if a && b {
		fmt.Printf("Condition is true\n")
	} else {
		fmt.Printf("Condition is not true\n")
	}

	//Called Logical NOT Operator. Use to reverses the logical state of its operand. If a condition is true, then Logical NOT operator will make it false.
	if !(a && b) {
		fmt.Printf("Condition is true\n")
	}

	/*
		OUTPUT:

		Condition is true
		Condition is not true
		Condition is true
	*/
}
