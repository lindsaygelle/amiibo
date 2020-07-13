package amiibo

import (
	"image"
)

type Image struct{}

func NewImage(URL string) (v Image, err error) {
	var i image.Image
	i, err = getRemoteImage(URL)
	return
}
