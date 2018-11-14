/*
	For more informations check:
		https://golang.org/pkg/strings/
*/

package main

import s "strings"
import "fmt"

var printIt = fmt.Println

func main() {
	//Create new strings
	var World = "world"
	var hello = "hello"

	// Concat strings.
	var helloWorld = hello + " " + World
	helloWorld += "!"
	printIt(helloWorld) // hello world!

	// functions available in strings
	printIt(len("Strings"))                                 // get length of string
	printIt("Contains:  ", s.Contains("test", "es"))        //check if string contain substring
	printIt("Count:     ", s.Count("test", "t"))            // check for string "te" occurence in test
	printIt("HasPrefix: ", s.HasPrefix("test", "te"))       //check prefix te in test
	printIt("HasSuffix: ", s.HasSuffix("test", "st"))       //check suffix st in test
	printIt("Index:     ", s.Index("test", "e"))            //find position in string
	printIt("Join:      ", s.Join([]string{"a", "b"}, "-")) //concatenate array strings
	printIt("Repeat:    ", s.Repeat("a", 5))                //repeat string n times
	printIt("Replace:   ", s.Replace("foo", "o", "0", -1))  //replace o with 0
	printIt("Replace:   ", s.Replace("foo", "o", "0", 1))   //replace o with 0 just 1
	printIt("Split:     ", s.Split("a-b-c-d-e", "-"))       //split string in array by -
	printIt("ToLower:   ", s.ToLower("TEST"))
	printIt("ToUpper:   ", s.ToUpper("test"))

	/*OUTPUT:
	hello world!
	7
	Contains:   true
	Count:      2
	HasPrefix:  true
	HasSuffix:  true
	Index:      1
	Join:       a-b
	Repeat:     aaaaa
	Replace:    f00
	Replace:    f0o
	Split:      [a b c d e]
	ToLower:    test
	ToUpper:    TEST
	*/
}
