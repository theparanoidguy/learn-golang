//Declare package name as "main"
package main

//Import the "fmt" package
import "fmt"

//Shortand without data type declaration cannot be done outside a function
//The declaration below will cause an error

//number := 5

//Start main function block
func main() {
  /*Declare two floating point numbers to pass into the "addStandard" function
  The general syntax to declare a variable is:
  var variable_name data_type = value
  */
  var num1 float64 = 4.2
  var num2 float64 = 10.1

// Shorthand method to delcare multiple variables of the same type
  /*In the below declaration, the numbers are said to be unpacked into
  the variables i.e. 10.1 is unpacked into "num4" & 4.2 is unpacked into
  "num3"*/
  var num3, num4 float64 = 4.2, 10.1

/*Shorthand dynamic declaration without specifying data ttpe. This can Only
be done inside a function i.e. inside "main" or any other function.
This is dynamic identification of data type. Go copiler will automatically
identify data types based on certain defaults. For example -
If  data type is not defined explicitly for a decimal point number, Go
identifies it as a "float64" by default. Thus, "num5" and "num6" below will
be identified as "float64"*/

  num5, num6 := 4.2, 10.1

  //Call function "addStandard"
    //Using "num1" & "num2"
  fmt.Println("The sum of two numbers is(Standard Declaration): ",addStandard(num1,num2))

    //Using "num3" & "num4"
  fmt.Println("The sum of two numbers is(Standard Declaration): ",addStandard(num3,num4))

    //Using "num5" & "num6"
  fmt.Println("The sum of two numbers is(Standard Declaration): ",addStandard(num5,num6))

    //Using "num1" & "num2"
    //Call function "addShorthandArgumentType"
  fmt.Println("The sum of two numbers is(Shorthand data type declaration): ",addShorthandArgumentType(num1,num2))

    //Using "num3" & "num4"
  fmt.Println("The sum of two numbers is(Shorthand data type declaration): ",addShorthandArgumentType(num3,num4))

    //Using "num5" & "num6"
  fmt.Println("The sum of two numbers is(Shorthand data type declaration): ",addShorthandArgumentType(num5,num6))

  //Call function "addStandardFloat32"
  
  /*The function below will fail. This is because as mentioned before, Go will identiy "num5" and "num6" as "float64"
  since the two variables are not defined explicitly. These "float64" variables are being passed into a function that
  accepts "float32". This is the cause of the error*/
  fmt.Println("The sum of two numbers is (Automatic data type identification): ",addStandardFloat32(num5,num6))
}

//Define a function called "addStandard" which takes in two numbers and adds them up
//and returns the result
//When we create a function that has values passed into it, we need to define
//the type(data type) of the values.
//"float64" is one such data type which represents a 64 bit floating point (decimal)
//number
//Additionally, we want the function to return the value of the operation.
//Hence, the return type also has to be defined.
//Since we are adding two 64 bit floating point numbers, its natural the result
//should be stored in the same data Type

//General format of a function that takes in values -
// func function_name(var1_name var1_type, var2_name var2_type) return_type {
//  statement1
//  statement2
//  return_statement
//}

/*We are going to declare various functions to achieve the same result
to demonstrate how it's done */
func addStandard (x float64, y float64) float64 {
  return x + y
}

/*Since "float64" is the common data type
for both x & y, we could write it like this -*/
func addShorthandArgumentType(x,y float64) float64 {
  return x + y
}

/*Same function taking "float32" numbers as inputs and returning "float32"
output*/
func addStandardFloat32(x, y float32) float32 {
  return x + y
}
