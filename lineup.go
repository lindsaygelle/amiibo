package amiibo

// https://www.nintendo.com/content/noa/en_US/amiibo/line-up/jcr:content/root/responsivegrid/lineup.model.json

// lineup is the unfettered Nintendo Amiibo lineup information provided by nintendo.com.
//
// lineup contains the product information for the Nintendo Amiibo product as well as some related metadata.
type lineup struct {

	// AmiiboList is a collection of Nintendo Amiibo products containing their product information.
	AmiiboList []lineupAmiibo `json:"amiiboList"`

	// ComponentPath is the relative path to the Nintendo resource file.
	ComponentPath string `json:"componentPath"`

	// Items is a collection of metadata related to Nintendo Amiibo products.
	Items []lineupItem `json:"item"`
}

// lineupAmiibo is the unfettered Nintendo Amiibo product information from nintendo.com.
//
// lineupAmiibo contains the English language product information for a specific Nintendo Amiibo.
type lineupAmiibo struct {

	// AmiiboName is the name of the Nintendo Amiibo product.
	//
	// AmiiboName can contain special characters that require filtering.
	AmiiboName string `json:"amiiboName"`

	// AmiiboPage is the relative URL to the Nintendo Amiibo product details page.
	//
	// AmiiboPage requires nintendo.com to be prepended to the AmiiboPage.
	AmiiboPage string `json:"amiiboPage"`

	// BoxArtURL is the relative URL to the Nintendo Amiibo box image page.
	//
	// BoxArtURL requires nintendo.com to be prepended to the URL.
	BoxArtURL string `json:"boxArtUrl"`

	// DetailsPath is the relative path to the Nintendo Amiibo product information page.
	DetailsPath string `json:"detailsPath"`

	// DetailsURL is the relative URL to the Nintendo Amiibo product information page.
	//
	// DetailsURL requires nintendo.com to be prepended to the URL.
	DetailsURL string `json:"detailsUrl"`

	// FigureURL is the relative URL to the Nintendo Amiibo product image.
	//
	// FigureURL requires nintendo.com to be prepended to the FigureURL.
	FigureURL string `json:"figureUrl"`

	// Franchise is the Nintendo or third-party product the Nintendo Amiibo product is affiliated with.
	//
	// Franchise is often null.
	Franchise string `json:"franchise"`

	// GameCode is the ID for the Nintendo game that the Nintendo Amiibo product is affiliated with.
	GameCode string `json:"gameCode"`

	// HexCode is the hexidecimal code for the Nintendo Amiibo product.
	HexCode string `json:"hexCode"`

	// OverviewDescription is the verbose Nintendo Amiibo product summary.
	//
	// OverviewDescription contains HTML tags and requires filtering.
	OverviewDescription string `json:"overviewDescription"`

	// Price is the price of the Nintendo Amiibo product in USD.
	//
	// Price need to be converted to a float.
	Price string `json:"price"`

	// PresentedBy is the namespace of the affiliated partner for the Nintendo Amiibo product.
	PresentedBy string `json:"presentedBy"`

	// ReleaseDateMask is the DD-MM-YYYY timestamp for the Nintendo Amiibo product release date.
	ReleaseDateMask string `json:"releaseDateMask"`

	// Series is the Nintendo product series the Nintendo Amiibo product is affiliated with.
	Series string `json:"series"`

	// Slug is the formatted namespace of the Nintendo Amiibo product.
	Slug string `json:"slug"`

	// Type is the Nintendo Amiibo product classification type.
	Type string `json:"type"`

	// UnixTimestamp is the Nintendo Amiibo product release date in milliseconds.
	UnixTimestamp int64 `json:"unixTimestamp"`

	// UPC is the universal product code for the Nintendo Amiibo product.
	UPC string `json:"upc"`
}

// lineupItem is the unfettered Nintendo Amiibo product additional information from nintendo.com.
//
// lineupItem contains additional information for a Nintendo Amiibo product.
type lineupItem struct {

	// Description is the verbose Nintendo Amiibo product summary.
	//
	// Description is often a null field.
	Description string `json:"description"`

	// LastModified is the Nintendo Amiibo product release date in milliseconds.
	LastModified int64 `json:"lastModified"`

	// Path is the relative path to the Nintendo Amiibo product information page according to the nintendo.com CDN.
	Path string `json:"path"`

	// Title is the name of the Nintendo Amiibo product.
	//
	// Title can contain special characters that require filtering.
	Title string `json:"title"`

	// URL is the relative path URL to the Nintendo Amiibo product information page.
	//
	// URL requires nintendo.com to be prepended to the URL.
	URL string `json:"url"`
}
