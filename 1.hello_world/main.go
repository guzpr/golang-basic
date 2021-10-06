//In all files, you must add the package declaration. Such declaration is
//composed of the keyword package then the name of the package. You should use
//main as package name for entry point file in your project

package main

// Import keyword is usually followed by an open parenthesis and a list of the
// program’s imported packages. Every package is written on a new line. Each
// package has a name delimited by double-quotes.
import (
	"fmt"
	"time"
)

//The main function is the entry point of the program. In every application, you
//have at least one main function.

func main() {
	fmt.Println("Hello world!", "2021")
	fmt.Println(time.Now())
}
