/*
	For more info check: https://golangbot.com/loops/
*/
package main

import "fmt"

func main() {
	//A for loop is a repetition control structure. It allows you to write a loop that needs to execute a specific number of times.
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}

	fmt.Println("===Break example===")
	//The break statement is used to terminate the for loop
	for a := 0; a < 10; a++ {
		if a == 5 {
			fmt.Println("Stop loop at 5")
			break //terminate loop
		}
		fmt.Printf("value of a: %d\n", a)
	}

	fmt.Println("===Continue example===")
	//The continue statement is used to skip the current iteration of the for loop
	for a := 0; a < 10; a++ {
		if a == 5 {
			fmt.Println("Skip 5")
			continue //terminate loop
		}
		fmt.Printf("value of a: %d\n", a)
	}

	/*
		OUTPUT:
			value of a: 0
			value of a: 1
			value of a: 2
			value of a: 3
			value of a: 4
			value of a: 5
			value of a: 6
			value of a: 7
			value of a: 8
			value of a: 9
			===Break example===
			value of a: 0
			value of a: 1
			value of a: 2
			value of a: 3
			value of a: 4
			Stop loop at 5
			===Continue example===
			value of a: 0
			value of a: 1
			value of a: 2
			value of a: 3
			value of a: 4
			Skip 5
			value of a: 6
			value of a: 7
			value of a: 8
			value of a: 9
	*/
}
