package amiibo

import (
	"path/filepath"
	"strings"
)

// ENGChartGame is the unfettered English language Nintendo Amiibo supported game information.
type ENGChartGame struct {

	// ID is the Nintendo game UUID.
	ID string `json:"id"`

	// Image is the relative path to the Nintendo game product image.
	Image string `json:"image"`

	// IsReleased is the availability status of the Nintendo game product.
	IsReleased string `json:"isReleased"`

	// Name is the unformatted name of the Nintendo game.
	Name string `json:"name"`

	// Path is the relative path to the Nintendo game product page.
	Path string `json:"path"`

	// ReleaseDateMask is the YYYY-MM-DD timestamp when the Nintendo game is or will be released.
	ReleaseDateMask string `json:"releaseDateMask"`

	// Type is the Nintendo product type.
	Type string `json:"type"`

	// URL is the relative path to the game page that supports a Nintendo Amiibo product.
	URL string `json:"url"`
}

// GetID returns the ID for ENGChartGame.
func (e ENGChartGame) GetID() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}

// ReadENGChartGame reads a ENGChartGame from disc.
func ReadENGChartGame(dir string, filename string) (v ENGChartGame, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteENGChartGame writes a ENGChartGame to disc.
func WriteENGChartGame(dir string, filename string, v *ENGChartGame) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
