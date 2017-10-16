//this class shows how struct type can be passed on to a function
package funcs
import "fmt"
type person struct{
	name string
	age int
}

func getpersonage(p person) int{
fmt.Println("inside get person's age")
return p.age

}

func getpersonname(p *person) string{
	fmt.Println("inside get person's name")
	return p.name
}
//Passstructtofunc works with struct type.
func Passstructtofunc(){
	fmt.Println("inside Passstructtofunc")
	p1:= person{name:"arun",age:30}
	p2:=person{name:"kumar",age:28}
	fmt.Println("get person's age:",getpersonage(p1))
	fmt.Println("get person's name:",getpersonname(&p2))
}