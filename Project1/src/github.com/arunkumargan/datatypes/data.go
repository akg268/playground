package data

/*Note if a variable is declared and not used it is an error in GO*/
import "fmt"

const constString = "I'm a constant string"

var outVar float32

/*WorkingWithDataTypes is a function to work with go lang data types */
func WorkingWithDataTypes() {
	fmt.Println("Inside Working with Data types function")
	fmt.Println(constString)
	outVar = 64.4568
	fmt.Println("global variable:", outVar)
	x := 10
	fmt.Println("value of x is:", x)
	fmt.Println("sum of 10+10 is:", 10+10)
	fmt.Println("Sting concatenation:" + "10+10")
	fmt.Println("bool false type", true && false)
	fmt.Println("bool true type", true || false)

	mystring := "This is a dynamic string"
	var typedstring string
	typedstring = "This is a typed sting"
	fmt.Println(mystring)
	fmt.Println(typedstring)
	//declare and assign in one line this is not required and go will infer the type by right hand side value
	var typedint int = 10
	fmt.Println("Typed int is:", typedint)
	//multiple var declaration in single line
	var a, b = 10, 20
	fmt.Println("a, b is :", a, b)
	var y, z = 10, "arun"
	fmt.Println("y, z is set to:", y, z)

}
