package font_test

import (
	"fmt"

	"github.com/xuender/font"
)

// ExampleDefaultFontPath is an example function.
func ExampleDefaultFontPath() {
	_, has := font.DefaultFontPath()
	fmt.Println(has)

	// Output:
	// true
}

// ExampleDefaultFont is an example function.
func ExampleDefaultFont() {
	data := font.DefaultFont()
	fmt.Println(len(data) > 0)

	// Output:
	// true
}
