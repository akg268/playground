package data

import "fmt"

/*WorkwithSlices to check how slices work. They are similar to arrays but can be used some of the
utility functions in the inbuilt package*/
func WorkwithSlices() {
	slice1 := make([]string, 3)
	fmt.Println("empty slice:", slice1)
	slice1[0] = "NewJersey"
	slice1[1] = "NewYork"
	fmt.Println("slice1:", slice1)
	slice2 := append(slice1, "Virginia", "California")
	fmt.Println("slice2:", slice2)
	fmt.Println("length of slice1:", len(slice1))
	fmt.Println("length of slice2:", len(slice2))
	fmt.Println("cap of slice1:", cap(slice1))
	fmt.Println("cap of slice2:", cap(slice2))
	fmt.Println("elements in slice:", slice2[0:1])

	slice3 := make([]string, len(slice2))
	fmt.Println("len of slice3:", len(slice3))
	fmt.Println("cap of slice3:", cap(slice3))
	copy(slice3, slice2)
	fmt.Println("elements in slice3", slice3)

}
