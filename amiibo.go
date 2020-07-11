package amiibo

import (
	"time"

	"golang.org/x/text/language"
)

// Amiibo is a generic interface for a Nintendo Amiibo product.
//
// Amiibo provides a handler to get all common fields between the ENGAmiibo and JPNAmiibo.
type Amiibo interface {

	// GetAvailable returns the Amiibo's availability status.
	//
	// GetAvailable is calculated by evaluating the Amiibo's release date is
	// before the current timestamp.
	GetAvailable() bool

	// GetID returns the Amiibo's hash key ID.
	GetID() string

	// GetLanguage returns the Amiibo's language identifier.
	GetLanguage() language.Tag
	GetMD5() (string, []byte, error)
	GetName() string
	GetNameAlternative() string
	GetPrice() string
	GetReleaseDate() time.Time
	GetSeries() string
	GetURL() string
}
