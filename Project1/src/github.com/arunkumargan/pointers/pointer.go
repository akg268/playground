package point

import "fmt"

//PointerFunc tries to change values of a variable by accessing its memory location

func PointerFunc(){
	fmt.Println("Inside pointer function")
a :=5
fmt.Println("a's value before calling pass by value:", a)
passbyvalue(a)
fmt.Println("a's value after calling pass by value:", a)

fmt.Println("a's value before calling pass by ref:", a)
passbyreference(&a)
fmt.Println("a's value after calling pass by ref:", a)
}

func passbyvalue(a int){
	a =10
fmt.Println("value of a inside passby value", a)
}

func passbyreference(a *int) {
*a=100
fmt.Println("value of a inside passby reference", *a)
}