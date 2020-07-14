package amiibo

import (
	"time"

	"golang.org/x/text/language"
)

// Software is a generic interface for an Nintendo software product.
//
// Software provides a handler to get all common fields between the ENGGame and JPNSoftware.
type Software interface {

	// GetAvailable returns the Nintendo software product availability status.
	//
	// GetAvailable is calculated by evaluating the Nintendo softwares products release date is
	// before the now's timestamp.
	GetAvailable() bool
	GetID() string
	GetLanguage() language.Tag
	GetMD5() (string, []byte, error)
	GetName() string
	GetNameAlternative() string
	GetReleaseDate() time.Time
	GetURL() string
}
