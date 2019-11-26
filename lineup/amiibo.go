package lineup

import (
	"path/filepath"
	"strings"
)

// Amiibo is a representation of a Nintendo Amiibo product provided from resource.Lineup.
// Amiibo contains data provided as-is from Nintendo with a mixture of content
// provided for each Nintendo Amiibo product to describe its unique attributes.
// Amiibo provided from lineup.Amiibo focus on describing the
// figurines meta data properties. Amiibos
// collected from the lineup resource are consumed by the amiibo/amiibo
// package to construct a normalized aggregation of an Amiibo across all resources.
type Amiibo struct {
	AmiiboName          string `json:"amiiboName"`
	AmiiboPage          string `json:"amiiboPage"`
	BoxArtURL           string `json:"boxArtUrl"`
	DetailsPath         string `json:"detailsPath"`
	DetailsURL          string `json:"detailsUrl"`
	FigureURL           string `json:"figureUrl"`
	Franchise           string `json:"franchise"`
	GameCode            string `json:"gameCode"`
	HexCode             string `json:"hexCode"`
	IsReleased          bool   `json:"isReleased"`
	OverviewDescription string `json:"overviewDescription"`
	PresentedBy         string `json:"presentedBy"`
	Price               string `json:"price"`
	ReleaseDateMask     string `json:"releaseDateMask"`
	Series              string `json:"series"`
	Slug                string `json:"slug"`
	Type                string `json:"type"`
	UnixTimestamp       int64  `json:"unixTimestamp"`
	UPC                 string `json:"upc"`
}

// Key returns a reliable ID.
func (a *Amiibo) Key() string {
	return strings.TrimSuffix(filepath.Base(a.DetailsURL), ".html")
}
