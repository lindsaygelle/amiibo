package image

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/network"
)

const (
	extGIF string = "GIF" // GIF file extension pattern.
	extJPG string = "JPG" // JPG file extension pattern.
	extPNG string = "PNG" // PNG file extension pattern.
)

const (
	measurement string = "px"
)

const (
	sep string = "." // sep string for parsing raw url
	rep string = ""  // rep string for parsing raw url
)

const (
	// Version is the semver of image.Image.
	Version string = "1.0.0"
)

// Image is a destructured image.Image provided by Go. Images are
// to give verbosity to all data pulled from the various Nintendo web resources
// that populate the content of the amiibo package.
type Image struct {
	Dir         string           `json:"dir"`
	Empty       bool             `json:"empty"`
	Ext         string           `json:"ext"`
	Height      int              `json:"height"`
	Measurement string           `json:"measurement"`
	Name        string           `json:"name"`
	Status      string           `json:"status"`
	StatusCode  int              `json:"status_code"`
	URL         *address.Address `json:"url"`
	Width       int              `json:"width"`
}

func (i *Image) String() string {
	return fmt.Sprintf("%s.%s", i.Name, i.Ext)
}

// NewImage creates a new instance of the image.Image based on the
// argument raw url string provided to the function. Returns an error
// if the argument raw url does not contain a http(s)://(subdomain|www) prefix
// or if url.Parse(rawurl) cannot parse the raw url. Will always return
// an instance of image.Image even if the raw url does not retun http.StatusOK.
// All image.Image's are created in reference to a remote Nintendo image resource.
func NewImage(rawurl string) (*Image, error) {
	var (
		bounds     image.Rectangle
		err        error
		height     = -1
		i          Image
		img        image.Image
		req, _     = http.NewRequest(http.MethodGet, rawurl, nil)
		status     = http.StatusText(http.StatusBadRequest)
		statusCode = http.StatusBadRequest
		width      = -1
		URL, _     = address.NewAddress(rawurl)
	)
	var (
		res, _ = network.Client.Do(req)
	)
	if res.StatusCode != statusCode {
		status = res.Status
		statusCode = res.StatusCode
	}
	defer res.Body.Close()
	i = Image{
		Dir:         parseDir(rawurl),
		Ext:         parseExt(rawurl),
		Height:      height,
		Measurement: measurement,
		Name:        parseName(rawurl),
		Status:      status,
		StatusCode:  statusCode,
		Width:       width,
		URL:         URL}
	img, err = parseImage(i.Ext, res.Body)
	if err != nil {
		return &i, err
	}
	bounds = img.Bounds()
	i.Empty = bounds.Empty()
	i.Height = bounds.Dy()
	i.Width = bounds.Dx()
	return &i, err
}

// parseDir parses the folder directory hosting the image file from the raw url string.
func parseDir(rawurl string) string {
	return filepath.Base(filepath.Dir(rawurl))
}

// parseExt parses the image file extension substring from the raw url string.
func parseExt(rawurl string) string {
	return strings.TrimPrefix(filepath.Ext(rawurl), sep)
}

// parseImage parses the image content body with the appopriate image decoder.
func parseImage(ext string, r io.ReadCloser) (image.Image, error) {
	var img image.Image
	var err error
	switch strings.ToUpper(ext) {
	case extGIF:
		img, err = gif.Decode(r)
	case extPNG:
		img, err = png.Decode(r)
	default:
		img, _, err = image.Decode(r)
	}
	return img, err
}

// parseName parses the image file name substring from the raw url string.
func parseName(rawurl string) string {
	return strings.Split(filepath.Base(rawurl), sep)[0]
}
