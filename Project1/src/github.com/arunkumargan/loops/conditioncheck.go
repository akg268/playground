package loops

import "fmt"
import "strings"

/*Checkcondition function prints the text based on choice of input. This is to show if else  */
func Checkcondition(input string) {
	//using == for string equality check expects exact match. To do case insensitive search lets use Strings package

	if strings.EqualFold(input, "fish") {
		fmt.Println("I like grilled fish")
	} else if strings.EqualFold(input, "chicken") {
		fmt.Println("I like grilled chicken")
	} else {
		fmt.Println("I like veggies")
	}
}
