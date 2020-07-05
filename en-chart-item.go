package amiibo

import (
	"path/filepath"
	"strings"
)

// ENGChartItem is the unfettered English language Nintendo Amiibo additional product information.
type ENGChartItem struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}

// GetID returns a ID for the ENGChartItem.
func (e ENGChartItem) GetID() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}
