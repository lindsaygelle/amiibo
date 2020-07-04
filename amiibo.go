package amiibo

// amiibo is the normalized amiibo data scraped from a rawAmiibo.
type amiibo struct{}

// amiiboCompatibility is the unfettered Nintendo Amiibo compatibility information provided by nintendo.com.
//
// amiiboCompatibility contains the relationship information between Nintendo Amiibo products
// and the games or applications that can be used within.
type amiiboCompatibility struct {
	AuthorMode    bool        `json:"authorMode"`
	AmiiboList    []rawAmiibo `json:"amiiboList"`
	ComponentPath string      `json:"componentPath"`
	GameList      []rawGame   `json:"gameList"`
	Items         []rawItem   `json:"items"`
	Language      string      `json:"language"`
	Mode          string      `json:"mode"`
}

// rawAmiibo is the unfettered Nintendo Amiibo product data provided by nintendo.com.
//
// rawAmiibo describes the abbreviated compatibility information for a specific Nintendo Amiibo figurine or card.
//
// rawAmiibo contains varying levels of completeness relative to the release status of the product.
type rawAmiibo struct{}

// rawGame is the unfettered game information related to a Nintendo Amiibo product provided by nintendo.com.
//
// rawGame describes the abbreviated game product information that has known Nintendo Amiibo support.
//
// rawGame contains varying levels of accuracy relative to the release status of Nintendo Amiibo products.
type rawGame struct{}

// rawItem is the unfettered auxiliary information related to a Nintendo Amiibo product provided by nintendo.com.
//
// rawItem describes the additional miscellaneous information that relates to Nintendo games that supports a Nintendo Amiibo product.
//
// rawItem contains varying levels of completeness relative to the release status of Nintendo Amiibo products or game titles.
type rawItem struct {

	// Description is the description for the Nintendo product.
	//
	// Description can commonly be a null field.
	Description string `json:"description"`
	// LastModified is the timestamp in milliseconds.
	LastModified int64 `json:"lastModified"`
	// Path is the relative path to the Nintendo game item to the Nintendo CDN.
	//
	// Path requires the nintendo.com domain prepended to the path.
	Path string `json:"path"`
	// Title is the name given to the Nintendo game item.
	//
	// Title often can contain special characters that will need filtering to prevent poor hashing keys.
	Title string `json:"title"`
	// URL is the relative URL to the Nintendo game item.
	//
	// URL requires the nintendo.com domain prepended to the URL.
	URL string `json:"url"`
}
