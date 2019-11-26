package game

import (
	"time"

	"golang.org/x/text/language"

	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/image"
)

const (
	// Version is the semver of game.Game.
	Version string = "1.0.0"
)

// Game is a structured representation of a Nintendo video-game that is compatible with a
// Nintendo Amiibo figuring product. Game structs are built from a mixture of resources that
// are provided from the amiibo/mix package.
// Games are consumed by amiibo/mux to create a basic HTTP REST API.
type Game struct {
	Compatability   []*Amiibo        `json:"compatability"`
	Complete        bool             `json:"complete"`
	Description     string           `json:"description"`
	GamePath        string           `json:"game_path"`
	GameURL         *address.Address `json:"game_url"`
	ID              string           `json:"id"`
	Image           *image.Image     `json:"image"`
	IsReleased      bool             `json:"is_released"`
	Language        language.Tag     `json:"language"`
	LastModified    int64            `json:"last_modified"`
	Path            string           `json:"path"`
	Name            string           `json:"name"`
	ReleaseDateMask string           `json:"release_date_mask"`
	Timestamp       time.Time        `json:"timestamp"`
	Title           string           `json:"title"`
	Type            string           `json:"type"`
	Unix            int64            `json:"unix"`
	URI             string           `json:"uri"`
	URL             *address.Address `json:"url"`
	Version         string           `json:"version"`
}
