package lineup

// Amiibo is a representation of a Nintendo Amiibo product provided from https://www.nintendo.com/amiibo/line-up/.
// Amiibo contains data provided as-is from Nintendo with a mixture of content
// provided for each Nintendo Amiibo product. Amiibo provided from lineup.Amiibo describes the
// Amiibo figure product data.
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
