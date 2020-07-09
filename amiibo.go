package amiibo

import (
	"time"

	"golang.org/x/text/language"
)

// Amiibo is a generic interface for an Amiibo product.
type Amiibo interface {
	GetID() string
	GetLanguage() language.Tag
	GetName() string
	GetNameAlternative() string
	GetPrice() string
	GetReleaseDate() time.Time
	GetSeries() string
	GetURL() string
}
