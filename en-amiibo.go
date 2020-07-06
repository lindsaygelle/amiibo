package amiibo

// ENGAmiibo is a formatted ENGChartAmiibo, ENGLineupAmiibo and ENGLineupItem.
type ENGAmiibo struct {
	Affiliation            string `json:"affiliation"`
	Availiable             string `json:"availiable"`
	BoxImage               string `json:"box_image"`
	Description            string `json:"description"`
	DetailsPath            string `json:"details_path"`
	DetailsURL             string `json:"details_url"`
	Epoch                  int64  `json:"epoch"`
	FigureImage            string `json:"figure_image"`
	Franchise              string `json:"franchise"`
	GameID                 string `json:"game_id"`
	Hex                    string `json:"hex"`
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	NameAlternative        string `json:"name_alternative"`
	Price                  string `json:"price"`
	Producer               string `json:"producer"`
	Product                string `json:"product"`
	ProductAlternative     string `json:"product_alternative"`
	ProductImage           string `json:"product_image"`
	ProductPage            string `json:"product_page"`
	ReleaseDate            string `json:"release_date"`
	ReleaseDateAlternative string `json:"release_date_alternative"`
	Series                 string `json:"series"`
	Title                  string `json:"title"`
	UPC                    string `json:"upc"`
	URL                    string `json:"url"`
	UUID                   string `json:"uuid"`
}

func (e ENGAmiibo) AddENChartAmiibo(v ENGChartAmiibo) {

	e.Affiliation = v.IsRelatedTo
	e.Availiable = v.IsReleased
	e.FigureImage = v.Image
	e.ID = v.TagID
	e.Name = v.Name
	e.ProductAlternative = v.Type
	e.ReleaseDateAlternative = v.ReleaseDateMask
	e.URL = v.URL
	e.UUID = v.ID
}

func (e ENGAmiibo) AddENLineupAmiibo(v ENGLineupAmiibo) {

	e.BoxImage = v.BoxArtURL
	e.Description = v.OverviewDescription
	e.DetailsPath = v.DetailsPath
	e.DetailsURL = v.DetailsURL
	e.Epoch = v.UnixTimestamp
	e.Franchise = v.Franchise
	e.GameID = v.GameCode
	e.Hex = v.HexCode
	e.NameAlternative = v.AmiiboName
	e.Price = v.Price
	e.Product = v.Type
	e.Producer = v.PresentedBy
	e.ProductImage = v.FigureURL
	e.ProductPage = v.AmiiboPage
	e.ReleaseDate = v.ReleaseDateMask
	e.Series = v.Series
	e.Title = v.Slug
	e.UPC = v.UPC
}
