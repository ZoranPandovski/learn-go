package main

import "fmt"

const (
	//iota identifier is used in const declarations to simplify definitions of incrementing numbers.
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 42
	fmt.Printf("%d\t%b\n", x, x)

	//Binary Left Shift Operator. The left operands value is moved left by the number of bits specified by the right operand.
	y := x << 1
	fmt.Printf("%d\t%b\n", y, y)

	//Binary Right Shift Operator. The left operands value is moved right by the number of bits specified by the right operand.
	z := x >> 1
	fmt.Printf("%d\t%b\n", z, z)

	fmt.Printf("%d\t%b\n", kb, kb)
	fmt.Printf("%d\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

	/*
		OUTPUT:
		42      101010
		84      1010100
		21      10101

		024    10000000000
		1048576 100000000000000000000
		1073741824      1000000000000000000000000000000
	*/
}
