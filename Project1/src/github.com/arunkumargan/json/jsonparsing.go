package jsn
import "encoding/json"
import "fmt"
//Note the fields in the Person structs should start with capital letters inorder for JSON to marshal the struct array completely
type Person struct{
	Name string
	Country string
}
type Persons[] Person

//Workwithjson function parses json
func Workwithjson(){
	var persons = Persons{
		Person{"A","US"},
		Person{"B","UK"},
	}

pj,_ :=	json.Marshal(persons)
fmt.Println(string(pj))
}

