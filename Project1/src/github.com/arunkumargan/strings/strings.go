package str

import s "strings"
import out "fmt"

//WorkingwithStringFunctions function to show go lang in built string functions
func WorkingwithStringFunctions(){
strin := "arun"
out.Println(s.ToUpper(strin))
out.Println(s.ToLower("ARUN"))
out.Println(s.Join([]string{"a","b"}, "*"))
out.Println(s.Index(strin,"u"))
out.Println(s.Split("a,b,c,d",","))
}