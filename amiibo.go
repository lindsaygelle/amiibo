package amiibo

import "time"

// Amiibo is a generic interface for an Amiibo product.
type Amiibo interface {
	GetID() string
	GetName() string
	GetNameAlternative() string
	GetPrice() string
	GetReleaseDate() time.Time
	GetSeries() string
	GetURL() string
}
