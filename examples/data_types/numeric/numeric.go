/*
	For more informations check:
		https://tour.golang.org/basics/11

*/
package main

import "fmt"

func main() {
	// Go arithmetic types represents a) integer types  b) floating point  c) other numeric types.

	/*Integers:
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	*/
	//Always use int unless you have a specific reason to use a sized or unsigned integer type.
	height := 450

	/*Float:
	float32 float64
	complex64 complex128
	*/
	width := 25.444

	/* Other
	byte rune
	*/

	//Print result:
	fmt.Println(width, height)
	// 25.444 450
}
