package testfmt

import (
	"fmt"
	"io"
	"os"
)

func SprintExample() {
	const name, age = "Kim", 22
	s := fmt.Sprint(name, " is ", age, " years old.\n")

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
