package inter
import "fmt"
/* interfaces are named collections of function signatures, to implement the interface the class must
provide implementation to all functions in the interface */
type Animal interface{

 eat() (string,string)
 run() (string,int)
}

type Animals struct{
	name string
	food string
	speed int
}


func (a Animals) eat() (string,string){
	fmt.Println(a.name)
	fmt.Println(a.eat)
	return a.name,a.food
}

func (a Animals) run() (string,int) {
	fmt.Println(a.name)
	fmt.Println(a.speed)
	return a.name,a.speed
}

func implementation(an Animal){
//fmt.Println(an)
fmt.Println(an.eat())
fmt.Println(an.run())
}



//Workwithinterrface function to work with interfaces
func Workwithinterface(){
	c:=Animals{"cheetah","deer",70}
	e:=Animals{"elephant","sugarcane",30}
   implementation(c)
   implementation(e)
}