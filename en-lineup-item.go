package amiibo

import (
	"path/filepath"
	"strings"
)

// ENGLineupItem is the unfettered Nintendo Amiibo product additional information from nintendo.com.
//
// ENGLineupItem contains additional information for a Nintendo Amiibo product.
type ENGLineupItem struct {

	// Description is the verbose Nintendo Amiibo product summary.
	//
	// Description is often a null field.
	Description string `json:"description"`

	// LastModified is the Nintendo Amiibo product release date in milliseconds.
	LastModified int64 `json:"lastModified"`

	// Path is the relative path to the Nintendo Amiibo product information page according to the nintendo.com CDN.
	Path string `json:"path"`

	// Title is the name of the Nintendo Amiibo product.
	//
	// Title can contain special characters that require filtering.
	Title string `json:"title"`

	// URL is the relative path URL to the Nintendo Amiibo product information page.
	//
	// URL requires nintendo.com to be prepended to the URL.
	URL string `json:"url"`
}

// GetID returns the ENGLineupItem ID.
func (e ENGLineupItem) GetID() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}
