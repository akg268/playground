package data

import "fmt"

/*WorkwithMaps function creates maps and manipulates it */
func WorkwithMaps() {

	map1 := make(map[string]string)
	map1["c1"] = "USA"
	map1["c2"] = "UK"
	map1["c3"] = "IN"
	v1 := map1["c1"]
	fmt.Println("value from map1:", v1)
	v2, k2 := map1["c2"]
	fmt.Println("value and is key present from map2:", v2, k2)
	delete(map1, "c3")
	fmt.Println("map1 after deleting:", map1)
}
