package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func TestImage(t *testing.T) {
	var _, err = amiibo.NewImage("https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png")
	if err != nil {
		t.Fatal(err)
	}
}
