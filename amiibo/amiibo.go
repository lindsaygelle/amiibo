package amiibo

import (
	"fmt"
	"reflect"
	"time"

	"golang.org/x/text/language"

	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/image"
)

const (
	// Version is the semver of amiibo.Amiibo.
	Version string = "1.0.0"
)

// Amiibo is a structured representation of a Nintendo Amiibo figuring.
// Amiibo structs are built from a mixture of resources that
// are provided from the amiibo/mix package.
// Amiibos are consumed by amiibo/mux to create a basic HTTP REST API.
type Amiibo struct {
	BoxImage        *image.Image     `json:"box_image"`
	Compatability   []*Game          `json:"compatability"`
	Complete        bool             `json:"complete"`
	Currency        string           `json:"currency"`
	Description     string           `json:"description"`
	DetailsPath     string           `json:"details_path"`
	DetailsURL      *address.Address `json:"details_url"`
	FigureURL       *address.Address `json:"figure_url"`
	Franchise       string           `json:"franchise"`
	FranchiseID     string           `json:"franchise_id"`
	GameCode        string           `json:"game_code"`
	HexCode         string           `json:"hex_code"`
	ID              string           `json:"id"`
	Image           *image.Image     `json:"image"`
	IsRelatedTo     string           `json:"is_related_to"`
	IsReleased      bool             `json:"is_released"`
	Language        language.Tag     `json:"language"`
	LastModified    int64            `json:"last_modified"`
	Name            string           `json:"name"`
	Overview        string           `json:"overview"`
	PageURL         *address.Address `json:"page"`
	Path            string           `json:"path"`
	PresentedBy     string           `json:"presented_by"`
	Price           string           `json:"price"`
	ReleaseDateMask string           `json:"release_date_mask"`
	Series          string           `json:"series"`
	SeriesID        string           `json:"series_id"`
	Slug            string           `json:"slug"`
	TagID           string           `json:"tag_id"`
	Timestamp       time.Time        `json:"timestamp"`
	Type            string           `json:"type"`
	TypeAlias       string           `json:"type_alias"`
	Unix            int64            `json:"unix"`
	UPC             string           `json:"upc"`
	URI             string           `json:"uri"`
	URL             *address.Address `json:"url"`
}

// Get gets a field from the Amiibo by its struct name and returns its string value.
func (a *Amiibo) Get(key string) string {
	var r = reflect.ValueOf(a)
	var v = reflect.Indirect(r).FieldByName(key)
	return fmt.Sprintf("%s", v)
}
