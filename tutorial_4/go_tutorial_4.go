/*
From here on, basic statements will not be commented elaborately. Only basic comments will be provided for statements like
"import" or "fmt.Println", etc.
New concepts and complex statements will follow the usual, elaborate commenting style.
*/

//Declare package name as "main"
package main

//Import the format package
import "fmt"

//Declare "main" function
func main() {

	/*
	   A pointer is a special type of data that holds the "position in memory" where the variable is stored.
	   This memory position is a hexadecial number like "0x19716daa" and corresponds to the actual memory cell in your RAM where
	   the value of the variable resides.
	   To illustrate
	*/

	//Declare variable
	var a int = 10
	var b int = 20
	var c float32 = 20.5
	d := 10.7

	/*
	   Declare pointer variable to store "address/location" of "a" in memory
	   The "&" operator before any variable name like "a" fetches it's memory address.
	   This address can be stored in a new variable. Thus, this new variable becomes the pointer of "a".
	   In the example below, "aAddr" is the pointer of "a" because it stores the memory address of "a".
	*/

	/*
		What would be the data type of a pointer variable? Without getting into the depth of it, we write it as

			*var_type [Where "var_type" can be "int", "float32", "float64", etc.]

		Thus, a pointer variable storing the address of an "int" has it's data type as "*int"
		A pointer variable storing the address of a "float32" has "*float32" as it's data type.
		The same rules of dynamic variables apply to pointers as well.
	*/
	var aAddr, bAddr *int = &a, &b
	var cAddr *float32 = &c

	//Example of dynamic allocation of pointer data types
	dAddr := &d

	fmt.Println("The address of variable 'a' is: ", aAddr)
	fmt.Println("The address of variable 'b' is: ", bAddr)
	fmt.Println("The address of variable 'c' is: ", cAddr)
	fmt.Println("The address of variable 'd' is: ", dAddr)

	/*
		Accesing pointer variables
		When we want to read the value stored at a memory location, we can use the "*" operator on the corresponding
		variable's pointer.
		An example is given below.
	*/

	/*
		In the statement below, we can see that we can use the operation "*" on "aAddr" to read through the memory addres
		stored in "aAddr" and fetch it's value
	*/
	fmt.Println("The memory address of variable 'a' is: ", aAddr, "and the value stored at this location is:", *aAddr)

	//We can perform mathemaical operaions on pointer references like this as well
	fmt.Println("To multiply 'a' & 'b' we can calculate the value of 'a * b' like so:")
	fmt.Println("a * b = ", a*b)
	fmt.Println("OR")
	fmt.Println("(*aAddr) * (*bAddr) = ", (*aAddr)*(*bAddr))
	fmt.Println("It makes no difference!")

	fmt.Println("We can change the value of 'a' by doing this: a = 20")
	a = 20
	fmt.Println("Value of 'a' is: ", a)
	fmt.Println("OR")
	*aAddr = 50
	fmt.Println("We can change it by directly modifying the data at the memory address like so: *aAddr = 50")
	fmt.Println("Value of 'a' now becomes: ", a)

	/*
		Pointers may seem like a new concept to programmers working on languages like NodeJS, Python or Java.
		However, it is one of the oldest programming concepts.
		Declaring and operating on pointers is always possible in low level languages like C/C++, Go, Rust, etc.
		As a matter of fact, it's quite impossible to write programs like system utilities, hardware drivers, etc. without
		using pointers.
	*/
}
