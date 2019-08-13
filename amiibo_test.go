package amiibo

import (
	"testing"
)

func TestAmiibo(t *testing.T) {

	rawAmiibo := &RawAmiibo{}

	newAmiibo(rawAmiibo)

	/*writeAmiibo(&Amiibo{Hex: "#ffdc81"})
	fmt.Println(getAmiibo("#ffdc81"))
	fmt.Println(deleteAmiibo(&Amiibo{Hex: "#ffdc81"}))*/
}
