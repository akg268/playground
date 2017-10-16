package loops
import "fmt"
import "time"
var z int
func Switchcase(x int){
  if(x >10){
   z =2
  }else{
    z=1
  }
switch z{
case 1:
fmt.Println("Number passed in is less than 10")
case 2:
fmt.Println("Number passed in is greater than 10")
default :
fmt.Println("Invalid input")
}
}
//Switch without condition
func SwitchcasebyTime(){
  t :=time.Now()
  fmt.Println(t.Weekday().String())
  switch t.Weekday().String(){
  case "Monday":
    fmt.Println("This is Monday")
  case "Wednesday":
  fmt.Println("This is Wednesday")
  default :
  fmt.Println("This is not Monday or Wednesday")
  }
}

//switch by type
func SwitchbyType(){

  myType := func(x interface{}){
  switch sw := x.(type){
  case bool:
    fmt.Println("this is a boolen")
  case int:
    fmt.Println("This is a int")
  default:
    fmt.Println("unknown type %T:", sw,"\n")
  }
}
  myType(false)
  myType(10)
  myType("string")
}
