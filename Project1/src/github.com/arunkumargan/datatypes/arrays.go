package data

import "fmt"
/*WorkwithArrays function to work on arrays*/
func WorkwithArrays(){
  fmt.Println("Inside working with Arrays")
  var lang[5] string
  lang[0]="JAVA"
  lang[1]="GO"
  lang[2]="NODE"
  lang[3]="JS"
  lang[4]="PYTHON"

  for i:=0;i<len(lang); i++{
    fmt.Println(lang[i])
  }
}
