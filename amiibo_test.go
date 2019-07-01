package amiibo

import (
	"fmt"
	"testing"
)

func TestAmiibo(t *testing.T) {

	writeAmiibo(&Amiibo{Hex: "#ffdc81"})
	fmt.Println(getAmiibo("#ffdc81"))
	fmt.Println(deleteAmiibo(&Amiibo{Hex: "#ffdc81"}))
}
