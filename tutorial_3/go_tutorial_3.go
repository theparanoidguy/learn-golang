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

/*Function declaration that returns multiple values
This function below would give an error. Even if your function returns two values which have the same data type, you need
to specify them.

	func returnStrings(x, y string) string {
		return x,y
	}

In the function above, even though the two values being returned are both strings, we cannot write the return type as string.
We'll need to write it as (string, string). The same applies to functions returning two or more int, float64, etc.
Correct declaration is given below
*/
func returnStrings(x, y string) (string, string) {
	return x, y
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

	//Declaring two strings
	var word1, word2 string = "Hello", "World!"

	//Dynamically declaring two strings
	word3, word4 := "Hello", "Again!"

	/*
		Type conversion
		Even though we cannot convert a dynamically allocated variable's data type after compilation, we can still do
		type conversion to use it as something else. For example, we're going to use "num7" & "num8" which are identified as
		"float64" by Go into a function that requires "float32". In order to do that, we'll convert the type of those variables.
		We'll use the two declared variables in the "addLongFloat32Format" function
	*/
	var num9 float32 = float32(num7)
	var num10 float32 = float32(num8)

	/*
		Type inference
		Let's consider a variable which is declared as such:
	*/
	var x float32 = 10.5
	y := x

	/*
		The value "10.5" is stored in the variable "y" dynamically. As we saw earlier, Go should assume the data type of such a value
		to be "float64". However, since "y" is actually storing  the value of "x" inside it, the type of "x" is considered while dynamically
		assigning the data type of "y". In this case, "y" becomes a "float32" variable.
		Proof of this is demonstrated in function call 13
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

	//Function Call - 9
	fmt.Println(returnStrings(word1, word2))

	//Function Call - 10
	fmt.Println(returnStrings(word3, word4))

	//Function Call - 11 - Using variables after converting their type and storing in a different variable
	fmt.Println("The sum of two numbers is(Variable type conversion from 'float64' to 'float32' w/ new variables): ", addLongFloat32Format(num9, num10))

	//Function call - 12 Using variables after converting their type at runtime but not storing them anywhere
	fmt.Println("The sum of two numbers is(Variable type conversion from 'float64' to 'float32' w/o new variables): ", addLongFloat32Format(float32(num7), float32(num8)))

	/*
		Function call - 13 Using a "float32" variable "x" and a dynamically assigned variable "y" that infers "x"
		If the variable "y" was not "float32", the compiler would have thrown an error at the function call below.
	*/
	fmt.Println("The sum of two numbers is(Using a 'float32' variable 'x' and another one dynamically infering 'x'): ", addLongFloat32Format(x, y))
}
