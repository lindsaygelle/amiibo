package amiibo

import (
	"path/filepath"
	"strings"
)

// ENGChartAmiibo is the unfettered English language Nintendo Amiibo product information.
type ENGChartAmiibo struct {

	// ID is the Nintendo Amiibo UUID.
	ID string `json:"id"`

	// Image is the relative path to the Nintendo Amiibo product image.
	Image string `json:"image"`

	// IsRelatedTo is the series or brand the Nintendo Amiibo product is related to.
	IsRelatedTo string `json:"isRelatedTo"`

	// IsReleased is the availability status of the Nintendo Amiibo product.
	IsReleased string `json:"isReleased"`

	// Name is the unformatted name of the Nintendo Amiibo product.
	Name string `json:"name"`

	// ReleaseDateMask is the YYYY-MM-DD timestamp when the Nintendo Amiibo product is or will be released.
	ReleaseDateMask string `json:"releaseDateMask"`

	// TagID is the ID assigned by Nintendo.
	TagID string `json:"tagid"`

	// Type is the Nintendo Amiibo product type.
	Type string `json:"type"`

	// URL is the relative path to the Nintendo Amiibo product page.
	URL string `json:"url"`
}

// GetID returns a ID for the ENGChartAmiibo.
func (e ENGChartAmiibo) GetID() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}

// ReadENGChartAmiibo reads a ENGChartAmiibo from disc.
func ReadENGChartAmiibo(dir string, filename string) (v ENGChartAmiibo, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteENGChartAmiibo writes a ENGChartAmiibo to disc.
func WriteENGChartAmiibo(dir string, filename string, v *ENGChartAmiibo) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
