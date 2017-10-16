package data

import "fmt"

/*WorkwithRanges is function to use range. Range is used in iterating over slices, arrays,maps and strings
range. Range gives index and value when iterating over slices and indexes
*/
func WorkwithRanges() {
	myarray := [3]int{1, 2, 3}
	sum := 0
	for _, elem := range myarray {
		sum += elem
	}
	fmt.Println("sum is :", sum)
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range map1 {
		fmt.Println("key and value from map is:", k, v)
	}

	for _, v := range map1 {
		fmt.Println("value from map is:", v)
	}
	for k, _ := range map1 {
		fmt.Println("key from map is:", k)
	}
}
