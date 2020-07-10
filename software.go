package amiibo

import (
	"time"

	"golang.org/x/text/language"
)

// Software is a generic interface for an Nintendo software product.
//
// Software provides a handler to get all common fields between the ENGGame and JPNSoftware.
type Software interface {
	GetAvailable() bool
	GetLanguage() language.Tag
	GetMD5() (string, error)
	Name() string
	NameAlternative() string
	ReleaseDate() time.Time
	GetURL() string
}
