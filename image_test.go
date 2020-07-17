package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func TestImage(t *testing.T) {
	var v, err = amiibo.GetImage("https://nintendo.com/content/dam/noa/en_US/amiibo/bowser-amiibo-super-smash-bros-series/screenshot-gallery/amiibo_Bowser_Smash_char.png")
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.WriteImage(filefolder, "bowser", &v)
	if err != nil {
		t.Fatal(err)
	}
	v, err = amiibo.ReadImage(filefolder, "bowser.png")
	if err != nil {
		t.Fatal(err)
	}
	if v.Ext != "PNG" {
		t.Fatalf("(Image).Ext %s != PNG", v.Ext)
	}
}
