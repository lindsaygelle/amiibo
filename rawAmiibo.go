package amiibo

import "fmt"

func newRawAmiibo() *RawAmiibo {
	return new(RawAmiibo)
}

var (
	_ rawAmiibo = (*RawAmiibo)(nil)
)

type rawAmiibo interface {
	String() string
}

// A RawAmiibo type represents a Nintendo Amiibo JSON object found in the raw Nintendo XHR HTTP response.
type RawAmiibo struct {
	AmiiboName          *RawAmiiboName        `json:"amiiboName"`          // "Toon Link - The Wind Waker"
	AmiiboPage          *RawAmiiboURL         `json:"amiiboPage"`          // "/amiibo/detail/toon-link…30th-anniversary-series"
	BoxArtURL           *RawAmiiboURL         `json:"boxArtUrl"`           // "/content/dam/noa/en_US/a…nk-WW_30thAnniv_box.png"
	DetailsPath         *RawAmiiboURL         `json:"detailsPath"`         // "/content/noa/en_US/amiib…30th-anniversary-series"
	DetailsURL          *RawAmiiboURL         `json:"detailsUrl"`          // "/amiibo/detail/toon-link…30th-anniversary-series"
	FigureURL           *RawAmiiboURL         `json:"figureURL"`           // "/content/dam/noa/en_US/a…k-WW_30thAnniv_char.png"
	Franchise           string                `json:"franchise"`           // "The Legend of Zelda"
	GameCode            string                `json:"gameCode"`            // "NVLEAK2A"
	HexCode             string                `json:"hexCode"`             // "#ffdc81"
	IsReleased          bool                  `json:"isReleased"`          // true
	OverviewDescription *RawAmiiboDescription `json:"overviewDescription"` // "<p style=\"margin-top: 0…ol the winds.\n</p>"
	PresentedBy         *RawAmiiboPresentedBy `json:"presentedBy"`         // "noa:publisher/nintendo"
	Price               *RawAmiiboPrice       `json:"price"`               // "24.99"
	ReleaseDateMask     *RawAmiiboReleaseDate `json:"releaseDateMask"`     // "12/02/2016"
	Series              string                `json:"series"`              // "30th anniversary"
	Slug                string                `json:"slug"`                // "toon-link-the-wind-waker…30th-anniversary-series"
	Type                string                `json:"type"`                // "Figure"
	UnixTimestamp       *RawAmiiboUnix        `json:"unixTimestamp"`       // 1480636800
	UPC                 string                `json:"upc"`                 // "045496893064"
}

func (r *RawAmiibo) Normalize() *Amiibo {
	return NewAmiibo(r)
}

func (r *RawAmiibo) String() string {
	return fmt.Sprintf("%v", r.AmiiboName)
}
