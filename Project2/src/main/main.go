package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// MyType is a struct
type Mytype struct {
	name string
	age  int
}

func (m Mytype) myfunc() {
	fmt.Println(m.name)
	fmt.Println(m.age)
}
func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println(time.Now())
	m := Mytype{"Arun", 33}
	fmt.Println("math.Abs(-11.2)", math.Abs(-11.2))
	fmt.Println("math.Ceil(10.3)", math.Ceil(10.3))
	fmt.Println("math.Copysign(1.2, -2.3)", math.Copysign(1.2, -2.3))
	fmt.Println("math.Dim(5.5, 4)", math.Dim(5.5, 4))
	fmt.Println("math.Gamma(1.1)", math.Gamma(1.1))
	fmt.Println("math.Max(11, 12)", math.Max(11, 12))
	fmt.Println("math.Log(11)", math.Log(11))
	fmt.Println("math.Log10(11)", math.Log10(11))
	fmt.Println("math.Min(10, 1)", math.Min(10, 1))
	fmt.Println("math.Nextafter(1, 2)", math.Nextafter(1, 2))
	fmt.Println("math.Pow(2, 4)", math.Pow(2, 4))
	fmt.Println("math.Sqrt(10)", math.Sqrt(10))
	fmt.Println("math.Trunc(11.2)", math.Trunc(11.2))
	var s = "Hello"
	fmt.Println("strings.ToUpper(s)", strings.ToUpper(s))
	fmt.Println("strings.NewReader(s).Size()", strings.NewReader(s).Size())
	fmt.Println("strings.Compare(a, b)", strings.Compare("b", "a"))
	m.myfunc()
	//functions.MyvoidFunction()
	funcs.MyvoidFunction()

}
