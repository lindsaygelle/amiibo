package amiibo

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Image is an image resource from Nintendo.
type Image struct {
	Ext    string      `json:"ext"`
	Image  image.Image `json:"-"`
	Height int         `json:"height"`
	Width  int         `json:"width"`
}

// GetImage returns a Image.
func GetImage(URL string) (v Image, err error) {
	var ext string
	var i image.Image
	i, ext, err = getRemoteImage(URL)
	if err != nil {
		return
	}
	var r = (i.Bounds().Max)
	v.Ext = ext
	v.Image = i
	v.Height = r.Y
	v.Width = r.Y
	return
}

// ReadImage reads an Image from disc.
func ReadImage(dir string, filename string) (v Image, err error) {
	var f *os.File
	var i image.Image
	f, err = os.Open(filepath.Join(dir, filename))
	if err != nil {
		return
	}
	i, _, err = image.Decode(f)
	if err != nil {
		return
	}
	var r = (i.Bounds().Max)
	v.Ext = strings.TrimPrefix(filepath.Ext(filename), ".")
	v.Image = i
	v.Height = r.Y
	v.Width = r.Y
	return
}

// WriteImage writes an Image to disc.
func WriteImage(dir string, filename string, v *Image) (fullpath string, err error) {
	var f *os.File
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	fullpath = filepath.Join(dir, fmt.Sprintf("%s.%s", filename, strings.ToLower(v.Ext)))
	f, err = os.Create(fullpath)
	if err != nil {
		return
	}
	defer f.Close()
	switch v.Ext {
	case "GIF":
		err = gif.Encode(f, v.Image, nil)
	case "JPG":
		err = jpeg.Encode(f, v.Image, nil)
	case "PNG":
		err = png.Encode(f, v.Image)
	default:
		err = fmt.Errorf("unsupported image format")
	}
	return
}
