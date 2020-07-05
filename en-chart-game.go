package amiibo

import (
	"path/filepath"
	"strings"
)

// ENGChartGame is the unfettered English language Nintendo Amiibo supported game information.
type ENGChartGame struct {
	Image           string `json:"image"`
	ID              string `json:"id"`
	IsReleased      string `json:"isReleased"`
	Name            string `json:"name"`
	Path            string `json:"path"`
	ReleaseDateMask string `json:"releaseDateMask"`
	Type            string `json:"type"`
	URL             string `json:"url"`
}

func (e ENGChartGame) String() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}
