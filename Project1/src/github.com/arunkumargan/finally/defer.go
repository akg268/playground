package finally

import "os"
import "fmt"

func createfile(path string) *os.File{

f,err := os.Create(path)
if err != nil{
	//panic(err)
	fmt.Println(err)
}
return f

}

func writefile(f *os.File) {

	fmt.Println("writing")
	fmt.Fprintln(f,"this is the data") 
}

func closefile(f *os.File){
 fmt.Println("closing file")
 f.Close()
}
//Workwithdefer function creates and writes a file and runs the defer function always
func Workwithdefer(){
 file:= createfile("/tmp/gotmpfile.txt")
writefile(file)
defer closefile(file)
}