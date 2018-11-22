/*

 */
package main

import "fmt"

//Go supports recursive functions. Here’s a classic factorial example.
//For more information check: https://www.golang-book.com/books/intro/7#section5
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))
}
