//Declare package name as "main"
package main

//Import required standard library packages
//You could import packages like so -
// import "fmt"
// import "math"
//However, this is not the recommended way.
//The recommended way is as follows

import ("fmt"
        "math"
        "math/rand")
//In the statement above, "rand" is a package inside "math" which we will use
//to generate random numbers

//Every single subpackage has to be imported seperately. "math" needs to be
//imported even if we have imported "math/rand" earlier

//The entire path up to the package that you want to use needs to be
//imported in Go

//There are no strict rules for indentation while importing, but it's a good
//practice to keep equal indentation at equal levels to make the code readable

//Declare main function
func main() {
  //The "Sqrt" is a function from the "math" package that calculates the
  //square root of the argument
  fmt.Println("The square root of 4 is: ",math.Sqrt(4))
  //sqrt_16()
  generateRandomNoSeed ()
}
func generateRandomNoSeed(){
  //Generates a random number using the "rand" package from "math"
  //The function to be called has to be written alone i.e. do not write
  //math.rand or "math/rand" when calling the function. Only write "rand"
  //The argument given to "rand" generates a random number within the range of
  //0 and the argument
  //No seed is supplied here. Thus, if you run the program multiple times
  //the same random number will be generated

  fmt.Println("Generating a number within 1 to 100 (No Seed): ",rand.Intn(100))
}

func generateRandomSeed() {
  //
}
func sqrt_16() {
  fmt.Println("Running function - foo")
  fmt.Println("The square root of 16 is: ",math.Sqrt(16))
}
//The "main" function in Go is simillar to C/C++ "main" and init classes.
//"main" will run without being explicitly called. However, other fucntions need
//to be explicitly called in order to execute
//In the above example, when we run this code, the statement inside "main"
//will run but the statement inside "sqrt_16" will not because "sqrt_16" was not
//called
//If you want to run "sqrt_16", uncomment line number 21 within "main".


//In case you're confused about what a function does/what all it can do
//use the "godoc" feature.
//"godoc" is a command line utility that gives more information about a
//function from a package.
//General syntax is "godoc package_name function_name"
//Example - godoc fmt Println =====> Will give information about the "Println"
//function from the "fmt" package
//**godoc is a command line utility i.e. don't write it in the code. Type it
//in your command line interpreter - bash, cmd, powershell, etc.
