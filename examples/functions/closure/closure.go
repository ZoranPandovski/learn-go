package main

import "fmt"

/*
Go supports anonymous functions, which can form closures.
 A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables;
 For more info check: https://www.golang-book.com/books/intro/7#section4
*/
func adder() func(int) int {
	sum := 0
	//adder function returns a closure. Each closure is bound to its own sum variable.
	return func(x int) int {
		sum += x
		return sum
	}

}

func main() {
	num := adder()
	for i := 0; i < 5; i++ {
		fmt.Println(num(i))
	}
}

/*
OUTPUT:
0
1
3
6
10
*/
