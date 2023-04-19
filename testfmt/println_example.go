package testfmt

import (
	"fmt"
)

func PrintlnExample() {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old.")

	// It is conventional not to worry about any
	// error returned by Println.

}
