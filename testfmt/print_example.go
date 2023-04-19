package testfmt

import (
	"fmt"
)

func PrintExample() {
	const name, age = "Kim", 22
	fmt.Print(name, " is ", age, " years old.\n")

	// It is conventional not to worry about any
	// error returned by Print.

}
