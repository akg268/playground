/***Make sure the GOPATH variable is set to point to your custom project directory
otherwise go will by default scan under /Usershome/go/src.
Here I set the GOPATH env variable to point to the Root path Project1
***/

package main

/*Note: the packages are directories. Each directory should have the same package but can be in a different file */
import (
	"fmt"

	"github.com/arunkumargan/funcs"
	"github.com/arunkumargan/loops"
	//"os"
	"github.com/arunkumargan/finally"
)

/*function main is the entry point for the standalone GO program
Main calls other functions in a blocking fashing, i.e the call to the callLoops happens
only after exiting from the myFirstFunction
*/
func main() {
	fmt.Print("Hello world!!\n")
	/* calling the first function from another function */
	myFirstFunction()
	callLoops()
	workwithArrays()
	workwithfunctions()
	callpointerfunc()
	/* creating currency structs here */

	c1 := data.Curreny{}
	c1.CountryName = "UK"
	c1.CurrencyName = "EURO"
	c2 := data.Curreny{"INDIA", "Rupees"}
	fmt.Println(c1)
	fmt.Println(c2.CurrencyName)
	fmt.Println("currency c2:", c2)
	c3 := &c1
	c3.CurrencyName = "POUND"
	fmt.Println(c1)

	x := 10
	fmt.Println(&x)
	z := &x
	fmt.Println(*z)
	workwithinterfaces()
	testerror()
	workwiththreads()
	jsn.Workwithjson()
	srt.WorkwithCustomSort()
	srt.Workwithsort()
	srt.Workwithstructsort()
	//checkpanic()
	finally.Workwithdefer()
	str.WorkingwithStringFunctions()
}
func myFirstFunction() {
	fmt.Print("before calling working with data types function\n")
	data.WorkingWithDataTypes()
	fmt.Println("myFirstFunction ended")
}

func callLoops() {
	//calling for loops
	loops.Looping()
	//calling conditional loops
	loops.Checkcondition("Chicken")
	//calling switch condition
	loops.Switchcase(-1)
	loops.Switchcase(9)
	loops.Switchcase(12)
	loops.SwitchcasebyTime()
	//switch by type
	loops.SwitchbyType()
}

func workwithArrays() {
	data.WorkwithArrays()
	data.WorkwithSlices()
	data.WorkwithMaps()
	data.WorkwithRanges()
}

func workwithfunctions() {

	sum := funcs.Funcwithreturn(1, 2)
	fmt.Println("sum returned from the function is:", sum)
	add, div := funcs.FuncwithMultiReturn(3, 20)
	fmt.Println("result of add is:", add)
	fmt.Println("result of divison is:", div)
	out := funcs.FuncwithvariableInput("hello", "world")
	fmt.Println("retuned string from FuncwithvariableInput function is: ", out)
	funcs.InvokeClosurefunc()
	funcs.Passstructtofunc()

}

func callpointerfunc() {
	point.PointerFunc()
}

func workwithinterfaces() {
	inter.Workwithinterface()
}

func testerror() {
	num, err := err.Testfunc(10)
	if err != nil {
		fmt.Println("error occured:", err)
	} else {
		fmt.Println("number is:", num)
	}
}

func workwiththreads() {
	threads.Workwiththreads()
}

/* func checkpanic(){
	panic("error occured")
 _, err := os.Create("llso/ooss")
  if err !=nil{
	  panic(err)
  }
} */
