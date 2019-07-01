package amiibo

import "fmt"

var (
	_ rawAmiibo = (*RawAmiibo)(nil)
)

func deleteRawAmiibo() bool {
	return false
}

func getRawAmiibo() *RawAmiibo {
	return nil
}

func writeRawAmiibo(rawAmiibo *RawAmiibo) bool {
	return false
}

type rawAmiibo interface {
	String() string
}

type RawAmiibo struct {
	AmiiboName          string `json:"amiiboName"`          // "Toon Link - The Wind Waker"
	AmiiboPage          string `json:"amiiboPage"`          // "/amiibo/detail/toon-link…30th-anniversary-series"
	BoxArtURL           string `json:"boxArtUrl"`           // "/content/dam/noa/en_US/a…nk-WW_30thAnniv_box.png"
	DetailsPath         string `json:"detailsPath"`         // "/content/noa/en_US/amiib…30th-anniversary-series"
	DetailsURL          string `json:"detailsUrl"`          // "/amiibo/detail/toon-link…30th-anniversary-series"
	FigureURL           string `json:"figureURL"`           // "/content/dam/noa/en_US/a…k-WW_30thAnniv_char.png"
	Franchise           string `json:"franchise"`           // "The Legend of Zelda"
	GameCode            string `json:"gameCode"`            // "NVLEAK2A"
	HexCode             string `json:"hexCode"`             // "#ffdc81"
	IsReleased          bool   `json:"isReleased"`          // true
	OverviewDescription string `json:"overviewDescription"` // "<p style=\"margin-top: 0…ol the winds.\n</p>"
	PresentedBy         string `json:"presentedBy"`         // "noa:publisher/nintendo"
	Price               string `json:"price"`               // "24.99"
	ReleaseDateMask     string `json:"releaseDateMask"`     // "12/02/2016"
	Series              string `json:"series"`              // "30th anniversary"
	Slug                string `json:"slug"`                // "toon-link-the-wind-waker…30th-anniversary-series"
	Type                string `json:"type"`                // "Figure"
	UnixTimestamp       int64  `json:"unixTimestamp"`       // 1480636800
	UPC                 string `json:"upc"`                 // "045496893064"
}

func (pointer *RawAmiibo) String() string {
	return fmt.Sprintf("%s", pointer.AmiiboName)
}
