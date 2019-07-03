package amiibo

import (
	"fmt"
	"testing"
)

func TestAmiiboSlice(t *testing.T) {
	b, err := local()
	if err != nil {
		panic(err)
	}
	slice := getAmiiboSlice(b)
	slice.Each(func(i int, amiibo *Amiibo) {
		fmt.Println(amiibo.ID)
	})
}
