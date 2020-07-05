package amiibo

import (
	"path/filepath"
	"strings"
)

// ENGChartAmiibo is the unfettered English language Nintendo Amiibo product information.
type ENGChartAmiibo struct {
	ID              string `json:"id"`
	Image           string `json:"image"`
	IsRelatedTo     string `json:"isRelatedTo"`
	IsReleased      string `json:"isReleased"`
	Name            string `json:"name"`
	ReleaseDateMask string `json:"releaseDateMask"`
	TagID           string `json:"tagid"`
	Type            string `json:"type"`
	URL             string `json:"url"`
}

// GetID returns a ID for the ENGChartAmiibo.
func (e ENGChartAmiibo) GetID() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}
