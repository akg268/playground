package loops

import "fmt"

/*Looping function:For is the only looping in GO. But we have if/else logic branching statements but not while loops or do whle loops */
func Looping() {
	fmt.Println("inisde Looping")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	count := 0
	fmt.Println("one more way of looping in for")
	for true {
		if count == 5 {
			break
		}
		fmt.Println(count)
		count++
	}
	fmt.Println("Looping with continue prints only odd numbers")
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
