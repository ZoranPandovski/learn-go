package main

import "fmt"

/*
A method is like a function, but attached to a type (called as receiver).
The official guide states “A method is a function with a special receiver argument”. The receiver appears in between the func keyword and the method name.
*/

type rect struct {
	width, height int
}

//area method
func (r *rect) area() int {
	return r.width * r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
}

/* OUTPUT:
area: 50
*/
