package funcs

import "fmt"

/*Funcwithreturn returns a value from the the function*/
func Funcwithreturn(a int, b int) int {
	fmt.Println("adding", a, "and", b)
	return a + b
}

/*FuncwithMultiReturn returns multiple return values from the the function*/
func FuncwithMultiReturn(a int, b int) (int, int) {
	sum := a + b
	divison := b / a
	return sum, divison

}

/*FuncwithvariableInput takes variable length input element. Here it takes string array and concatinates it */
func FuncwithvariableInput(input ...string) string {
	var x string
	for _, v := range input {
		x += v
	}
	return x
}
