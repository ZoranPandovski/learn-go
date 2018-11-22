package main

import "fmt"

/*Go has built-in support for multiple return values.
This feature is used often in idiomatic Go, for example to return both result and error values from a function.
More info: https://gobyexample.com/multiple-return-values
*/

//swap returns 2 string
func swap(name1, name2 string) (string, string) {
	return name2, name1
}

func main() {
	//Multiple values are often used to return an error value along with the resul
	fmt.Println(swap("Name1", "Name2"))
}
