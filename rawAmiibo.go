package amiibo

import "fmt"

func newRawAmiibo() *RawAmiibo {
	return &RawAmiibo{}
}

func NewRawAmiibo() *RawAmiibo {
	return &RawAmiibo{}
}

var (
	_ rawAmiibo = (*RawAmiibo)(nil)
)

type rawAmiibo interface {
	String() string
}

// A RawAmiibo type represents a Nintendo Amiibo JSON object found in the raw Nintendo XHR HTTP response.
type RawAmiibo struct {
	AmiiboName          *RawAmiiboName        `json:"amiiboName"`
	AmiiboPage          *RawAmiiboURL         `json:"amiiboPage"`
	BoxArtURL           *RawAmiiboURL         `json:"boxArtUrl"`
	DetailsPath         *RawAmiiboURL         `json:"detailsPath"`
	DetailsURL          *RawAmiiboURL         `json:"detailsUrl"`
	FigureURL           *RawAmiiboURL         `json:"figureURL"`
	Franchise           string                `json:"franchise"`
	GameCode            string                `json:"gameCode"`
	HexCode             string                `json:"hexCode"`
	IsReleased          bool                  `json:"isReleased"`
	OverviewDescription *RawAmiiboDescription `json:"overviewDescription"`
	PresentedBy         *RawAmiiboPresentedBy `json:"presentedBy"`
	Price               *RawAmiiboPrice       `json:"price"`
	ReleaseDateMask     *RawAmiiboReleaseDate `json:"releaseDateMask"`
	Series              string                `json:"series"`
	Slug                string                `json:"slug"`
	Type                string                `json:"type"`
	UnixTimestamp       *RawAmiiboUnix        `json:"unixTimestamp"`
	UPC                 string                `json:"upc"`
}

func (r *RawAmiibo) String() string {
	return fmt.Sprintf("%v", r.AmiiboName)
}
