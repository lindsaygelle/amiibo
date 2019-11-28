package image_test

import (
	"fmt"
	"testing"

	"github.com/gellel/amiibo/image"
)

func Test(t *testing.T) {
	const (
		rawurl = "https://www.nintendo.com/content/dam/noa/en_US/amiibo/alm-amiibo-fire-emblem-series/screenshot-gallery/amiibo_Alm_FireEmblem_char.png"
	)
	var (
		image, err = image.NewImage(rawurl)
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(image)
}
