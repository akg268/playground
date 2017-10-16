package err
import "errors"
import "fmt"

//Testfunc checks the error, custom errors can be thrown by using errors package.
func Testfunc(a int) (int, error){
	fmt.Println("a value is:",a)
	if a == 10{
		return 0, errors.New("Not allowed error")
	}
	return a, nil
}

