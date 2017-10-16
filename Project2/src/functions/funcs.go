package funcs

import (
	"fmt"
)

//MyvoidFunction returns nothingit is an exported function can be accessed from another package.
func MyvoidFunction() {

	s := fmt.Sprint("Sprint", "prints", "variable", "interface", "type", "data", "and", "returns", "single", "string")
	fmt.Print(s)
}
