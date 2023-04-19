package testfmt

import (
	"fmt"
	"io"
	"os"
)

func SprintfExample() {
	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
