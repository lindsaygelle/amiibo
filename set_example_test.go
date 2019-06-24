package amiibo_test

import (
	"fmt"

	"github.com/gellel/amiibo"
)

func ExampleNewSet() {

	a := &amiibo.Amiibo{
		Head:  "00000000",
		Image: "https://raw.githubusercontent.com/N3evin/AmiiboAPI/master/images/icon_00000000-00000002.png",
		Name:  "Mario",
		Tail:  "00000002"}

	b := &amiibo.Amiibo{
		Head:  "00000000",
		Image: "https://raw.githubusercontent.com/N3evin/AmiiboAPI/master/images/icon_00000000-00340102.png",
		Name:  "Mario",
		Tail:  "00340102"}

	fmt.Println(amiibo.NewSet(a, b))
	// Output: &map[Mario:&[Mario Mario]]
}
