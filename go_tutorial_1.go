//Comments begin with "//". These are ignored by the compiler and are
//used for describing the written code or disabling/commenting out specific code
//without erasing it.

//Every Go program starts with a package declariation i.e. it is a part of a package
//For the most part, we can just create a "main" package and move on.
//For more comple projects, splitting programs into multiple packages may be necessary

//Declare package as "main"
package main

//Import the "format" package. In Go, single quotes are almost never used.
//Make a habit of using double quotes when writing Go code.
import "fmt"

//Declare a function using the basic syntax "func function_name(parameters) {}"
func main() {
  //"Println" function from the "fmt" package prints a piece of text and adds a new line after that.
  fmt.Println("Welcome to Go!")
}

//In order to define the body of a function, the curly braces must start on the same line as the
//function definition
//Example -
// func test1() {} ===============> OKAY! Will work.
// func test2() {  ===============> NOT OKAY! This will throw an error.
// }
