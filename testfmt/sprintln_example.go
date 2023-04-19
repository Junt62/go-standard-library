package testfmt

import (
	"fmt"
	"io"
	"os"
)

func SprintlnExample() {
	const name, age = "Kim", 22
	s := fmt.Sprintln(name, "is", age, "years old.")

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
