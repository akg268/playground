package threads
import (
	"fmt"
	"strconv"
)
//Workwiththreads function invokes localfunction in seperate threads. 
func Workwiththreads(){
	for range [5]int {1,2,3,4,5}{
localfunction("direct call")
	}
for i:=0; i<5; i++{
go localfunction("go routine:" + strconv.Itoa(i))
}
var consoleInput string
fmt.Scanln(&consoleInput)
fmt.Println("done")
}

func localfunction(name string){
fmt.Println(name)
}


