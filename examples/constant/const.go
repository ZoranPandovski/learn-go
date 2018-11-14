/*
Go supports constants of character, string, boolean, and numeric values.

Run code
package main
import "fmt"
import "math"
const declares a constant value.
*/

package main

import "fmt"

func main() {

	const Pi = 3.14
	const Success = true
	const LocalFIle string = "usr/lib/ssss"

	fmt.Println(Pi, Success, LocalFIle)
	//3.14 true usr/lib/ssss
}
