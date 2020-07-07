//Declare package name as "main"
package main

//Import the "fmt" package
import "fmt"

/*
Define functions which take in two numbers, add them up and returns the result. When we create a function that has values passed into it, we need to define
the type(data type) of the values. "float32" is one such data type which represents a 32 bit floating point (decimal)
number. Additionally, we want the function to return the value of the operation. Hence, the return type also has to be defined.
Since we are adding two 32 bit floating point numbers, its natural the result should be stored in the same data Type

General format of a function that takes in values -

	func function_name(var1_name var1_type, var2_name var2_type) return_type {
		statement1
	  	statement2
	  	return_statement
	}

We are going to declare various functions to achieve the same result
to demonstrate how it's done.
*/

//Function Declaration Format - 1 - Standard/Long Format
func addShortFloat32Format(x float32, y float32) float32 {
	return x + y
}

//Function Declaration Format - 2 - Short Format
func addLongFloat32Format(x, y float32) float32 {
	return x + y
}

//Start main function block
func main() {
	/*
	  Variable Declaration Format - 1 - Static Long Format
	  Declare two floating point numbers
	*/
	var num1 float32 = 4.2
	var num2 float32 = 10.1

	/*
	  Variable Declaration Format - 2 - Static Short Format
	  We can declare the same in using a shorter format.
	  When this is done, the values are assigned in the same sequence as variable names.
	*/
	var num3, num4 float32 = 4.2, 10.1

	/*
	  Variable Declaration Format - 3 - Dynamic Long Format
	  We can declare a variable without using a data type and only assigning a value.
	  In these cases, Go will automatically assign data types to the variable at compile time.
	  Please note that Go has certain defaults for variables. For example, for the below declaration
	  Go automatically assigns the type "float64" to this variable.
	  This type of declaration is done using the ":=" symbol.
	  The var keyword is skipped in this format.
	*/
	num5 := 4.2
	num6 := 10.1

	/*
	  Variable Declatation format - 4 - Dynamic Short Format
	  The same concept as format 2 applied to dynamic declarations.
	*/
	num7, num8 := 4.2, 10.1

	/*
		Please note that variable declaration types 3 and 4 shown above can only be done within a function.
		The function can be "main" or a user defined function. In this case, we are declaring them
		inside "main"
	*/

	/*
		Function Call - 1
		Using "num1" & "num2"
	*/
	fmt.Println("The sum of two numbers is(Static Long Variable Declaration, Short Function Declaration 'float32' i/o): ", addShortFloat32Format(num1, num2))

	/*
		Function Call - 2
		Using "num3" & "num4"
	*/
	fmt.Println("The sum of two numbers is(Static Short Variable Declaration, Short Function Declaration 'float32' i/o): ", addShortFloat32Format(num3, num4))

	/*
		Function Call - 3
		Using "num5" & "num6"
	*/
	//fmt.Println("The sum of two numbers is(Dynamic Long Variable Declaration, 'float32' i/o): ", addShortFloat32Format(num5, num6))

	/*
		Function Call - 4
		Using "num7" & "num8"
	*/
	//fmt.Println("The sum of two numbers is(Dynamic Long Variable Declaration, 'float32' i/o): ", addShortFloat32Format(num7, num8))

	/*
		Function Call - 5
		Using "num1" & "num2"
	*/
	fmt.Println("The sum of two numbers is(Static Long Variable Declaration, Long Funtion Declaration 'float32' i/o): ", addLongFloat32Format(num1, num2))

	/*
		Function Call - 6
		Using "num3" & "num4"
	*/
	fmt.Println("The sum of two numbers is(Static Short Variable Declaration, Long Function Declaration 'float32' i/o): ", addLongFloat32Format(num3, num4))

	/*
		Function Call - 7
		Using "num5" & "num6"
	*/
	//fmt.Println("The sum of two numbers is(Dynamic Long Variable Declaration, 'float32' i/o): ", addLongFloat32Format(num5, num6))

	/*
		Function Call - 8
		Using "num7" & "num8"
	*/
	//fmt.Println("The sum of two numbers is(Dynamic Long Variable Declaration, 'float32' i/o): ", addLongFloat32Format(num7, num8))

	/*
		Function calls 3, 4, 7, 8 have been commented out because they will cause errors while compiling.
		Do you remember we discussed earlier that when Go dynamically assigns data type to a variable, it cannot be changed later.
		The numbers "num5", "num6", "num7" & "num8" are decimal point numbers that were assigned data types dynamically as ""float64".
		We cannot pass them into a function that expects 'float32' input.
		If you wish to see the error message that will appear, enable the calls by removing the '//' comment before them and try to
		compile the program.
	*/

	/*
		Additioanlly, we have also written a statement below to simply print the values of the dynamically declared float64 variables.
		This is because in Go, we cannot declare a variable and not use it i.e. unused variables are forbidden. The compiler throws errors
		if unused variables are encountered in the code. Thus, we have used them in a print statement below.
		If you wish to see the error the compiler would throw, comment out the line of code below.
	*/
	fmt.Println("The value of 'num5', 'num6', 'num7' & 'num8' are:", num5, num6, num7, num8, "respectively")
}
