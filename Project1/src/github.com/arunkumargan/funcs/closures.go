package funcs

import "fmt"

func closurefunction() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

/*InvokeClosurefunc is to call the local function */
func InvokeClosurefunc() {
	y := closurefunction()
	for x := range [5] int {1,2,3,4,5}{
		fmt.Println("x value is:",x)
		fmt.Println(y())
	}
	z := closurefunction()
	fmt.Println(z())
}
