package amiibo_test

import (
	"fmt"

	"github.com/gellel/amiibo"
)

func NewMapExample() {

	ryu := &amiibo.Amiibo{Name: "Ryu"}
	ken := &amiibo.Amiibo{Name: "Ken"}

	fmt.Println(amiibo.NewMap(ryu, ken))
	// Output: &map[Ryu:Ryu Ken:Ken]
}
