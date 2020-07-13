package amiibo

import (
	"image"
)

// Image is an image resource from Nintendo.
type Image struct {
	Ext    string `json:"ext"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

func NewImage(URL string) (v Image, err error) {
	var i image.Image
	i, ext, err = getRemoteImage(URL)
	if err != nil {
		return
	}
	var r = i.Bounds().Max
	v.Ext = ext
	v.Height = r.Y
	v.Width = r.Y
	return
}
