package amiibo

import (
	"path/filepath"
	"strings"
)

// ENGChartItem is the unfettered English language Nintendo Amiibo additional product information.
type ENGChartItem struct {

	// Description is the Nintendo Amiibo item product description.
	Description string `json:"description"`

	// LastModified is the Nintendo Amiibo item product release date in milliseconds.
	LastModified int64 `json:"lastModified"`

	// Path is the relative path to the Nintendo Amiibo item product page.
	Path string `json:"path"`

	// Title is the unformatted name of the Nintendo Amiibo item product.
	Title string `json:"title"`

	// URL is the relative path to the Nintendo Amiibo item page.
	URL string `json:"url"`
}

// GetID returns a ID for the ENGChartItem.
func (e ENGChartItem) GetID() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}

// ReadENGChartItem reads a ENGChartItem from disc.
func ReadENGChartItem(dir string, filename string) (v ENGChartItem, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteENGChartItem writes a ENGChartItem to disc.
func WriteENGChartItem(dir string, filename string, v *ENGChartItem) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
