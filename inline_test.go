package inline_test

import (
	"fmt"

	"github.com/lmittmann/inline"
)

func ExampleIf() {
	isEnabled := true
	fmt.Println(inline.If(isEnabled, "enabled", "disabled"))
	// Output:
	// enabled
}

func ExamplePtr() {
	var v *int = inline.Ptr(42)
	fmt.Println(*v)
	// Output:
	// 42
}
