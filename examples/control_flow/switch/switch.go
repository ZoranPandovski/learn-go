/*
	For more info check: https://gobyexample.com/switch
*/
package main

import "fmt"

var printIt = fmt.Println

func main() {

	//A switch statement allows a variable to be tested for equality against a list of values.
	var grade string
	var marks int = 80

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		printIt("Excelent")
	case grade == "B":
		printIt("Well done")
	case grade == "C":
		printIt("Passed")
	case grade == "D":
		printIt("Please, try again")
	default:
		printIt("Invalid grade.")
	}

	/*
		OUTPUT:
		Well done
	*/
}
