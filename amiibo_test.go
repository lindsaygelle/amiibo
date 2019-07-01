package amiibo

import (
	"fmt"
	"testing"
)

func TestAmiibo(t *testing.T) {

	amiibo := getAmiibo()

	for _, amiibo := range amiibo {
		fmt.Println(amiibo.URL)
	}
}
