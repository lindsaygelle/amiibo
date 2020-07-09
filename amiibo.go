package amiibo

import (
	"time"

	"golang.org/x/text/language"
)

// Amiibo is a generic interface for an Amiibo product.
//
// Amiibo provides a handler to get all common fields between the ENGAmiibo and JPNAmiibo.
type Amiibo interface {
	GetID() string
	GetLanguage() language.Tag
	GetMD5() (string, error)
	GetName() string
	GetNameAlternative() string
	GetPrice() string
	GetReleaseDate() time.Time
	GetSeries() string
	GetURL() string
}
