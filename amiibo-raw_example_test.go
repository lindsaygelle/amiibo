package amiibo_test

import (
	"fmt"

	"github.com/gellel/amiibo"
)

func ExampleNewRawAmiibo() {
	var (
		character = "Fox"
		game      = "Star Fox"
		head      = "05800000"
		image     = "https//raw.githubusercontent.com/N3evin/AmiiboAPI/master/images/icon_05800000-00050002.png"
		name      = "Fox"
		series    = "Super Smash Bros."
		tail      = "00050002"
		t         = "Figure"
		AU        = "2014-11-29"
		EU        = "2014-11-28"
		JP        = "2014-12-06"
		NA        = "2014-11-21"
	)
	fmt.Println(amiibo.NewRawAmiibo(series, character, game, head, image, name, tail, t, AU, EU, JP, NA))
	// Output: 0580000000050002
}
