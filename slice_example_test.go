package amiibo_test

import (
	"fmt"

	"github.com/gellel/amiibo"
)

func ExampleNewSlice() {

	mario := &amiibo.Amiibo{Name: "Mario"}
	luigi := &amiibo.Amiibo{Name: "Luigi"}

	fmt.Println(amiibo.NewSlice(mario, luigi))
	// Output: &[Mario Luigi]
}
