package srt
import "fmt"
import "sort"

type Bylength []string

func (fr Bylength) Len() int{
	return len(fr)
}

func (fr Bylength) Swap(a , b int){
	fr[a], fr[b] = fr[b], fr[a]
}

func (fr Bylength) Less(a, b int) bool{
return len(fr[a]) < len(fr[b])
}
//WorkwithCustomSort function is to display implementation of  sort interface
func WorkwithCustomSort(){
	fruits :=[]string{"watermelon","muskmelon","honeydew"}
	sort.Sort(Bylength(fruits))
	fmt.Println(fruits)
}
//Workwithsort function sorts by go data types
func Workwithsort(){
	intarr := []int{2,4,1,6,8}
	sort.Ints(intarr)
	fmt.Println(intarr)
	strarr := []string{"o","h","l","e","l"}
	sort.Strings(strarr)
	fmt.Println(strarr)
}

type Person struct{
	Name string
	Age int
}

type Persons []Person

func (ps Persons) Len() int{
return	len(ps)
 }

func (ps Persons) Swap(i, j int){
	ps[i], ps[j] = ps[j], ps[i]
}
//sorting by age in asc order
func (ps Persons) Less(i, j int) bool{
	return ps[i].Age < ps[j].Age
}
func Workwithstructsort(){
 persons := Persons{
	 Person{"kumar",20},
	 Person{"arun",25},
 }
 
 sort.Sort(persons)
fmt.Println(persons)
}