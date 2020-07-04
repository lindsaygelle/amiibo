package amiibo

// amiibo is the normalized amiibo data scraped from a compatabilityAmiibo.
type amiibo struct{}

// amiiboCompatibility is the unfettered Nintendo Amiibo compatibility information provided by nintendo.com.
//
// amiiboCompatibility contains the relationship information between Nintendo Amiibo products
// and the games or applications that can be used within.
type amiiboCompatibility struct {
	AuthorMode    bool                  `json:"authorMode"`
	AmiiboList    []compatabilityAmiibo `json:"amiiboList"`
	ComponentPath string                `json:"componentPath"`
	GameList      []compatabilityGame   `json:"gameList"`
	Items         []compatabilityItem   `json:"items"`
	Language      string                `json:"language"`
	Mode          string                `json:"mode"`
}

// compatabilityAmiibo is the unfettered Nintendo Amiibo product data provided by nintendo.com.
//
// compatabilityAmiibo describes the abbreviated compatibility information for a specific Nintendo Amiibo figurine or card.
//
// compatabilityAmiibo contains varying levels of completeness relative to the release status of the product.
type compatabilityAmiibo struct {

	// compatabilityCommon is a composed property.
	compatabilityCommon

	// ID is the unique identifier for the Nintendo Amiibo figurine or card.
	//
	// ID is a UUID.
	ID string `json:"id"`

	// IsRelatedTo is the name of a Nintendo product or series that the Nintendo Amiibo figurine or card is affiliated with.
	//
	// IsRelatedTo often can contain special characters that will need filtering to prevent poor hashing keys.
	IsRelatedTo string `json:"isRelatedTo"`

	// Image is the relative URL to the Nintendo Amiibo box art image.
	//
	// Image requires the nintendo.com domain prepended to the URL.
	Image string `json:"image"`

	// Name is the name given to the Nintendo game.
	//
	// Name often can contain special characters that will need filtering to prevent poor hashing keys.
	Name string `json:"name"`

	// TagID is the identifier namespace for the Nintendo Amiibo figurine or card.
	TagID string `json:"tagId"`

	// Type is the entity type for the Nintendo Amiibo product.
	//
	// Type can be type amiibo or other for Nintendo Amiibo products.
	Type string `json:"type"`
}

// compatabilityCommon is the common properties shared between compatabilityAmiibo, compatabilityGame and compatabilityItem structs.
type compatabilityCommon struct {

	// Path is the relative path to the Nintendo product and its information in the context of the Nintendo CDN.
	//
	// Path requires the nintendo.com domain prepended to the path.
	Path string `json:"path"`

	// ReleaseDateMask is the YYYY-MM-DD expression for the date when the Nintendo product was released.
	ReleaseDateMask string `json:"releaseDateMask"`

	// URL is the relative URL to the Nintendo product.
	//
	// URL requires the nintendo.com domain prepended to the URL.
	URL string `json:"url"`
}

// compatabilityGame is the unfettered game information related to a Nintendo Amiibo product provided by nintendo.com.
//
// compatabilityGame describes the abbreviated game product information that has known Nintendo Amiibo support.
//
// compatabilityGame contains varying levels of accuracy relative to the release status of Nintendo Amiibo products.
type compatabilityGame struct {

	// compatabilityCommon is a composed property.
	compatabilityCommon

	// ID is the unique identifier for the Nintendo game.
	//
	// ID is a UUID.
	ID string `json:"id"`

	// Image is the relative URL to the Nintendo game box art image.
	//
	// Image requires the nintendo.com domain prepended to the URL.
	Image string `json:"image"`

	// IsReleased is a string representation of a boolean that indicates whether the Nintendo game is available.
	//
	// IsReleased needs to be formatted to a bool data type.
	IsReleased string `json:"isReleased"`

	// Name is the name given to the Nintendo game.
	//
	// Name often can contain special characters that will need filtering to prevent poor hashing keys.
	Name string `json:"name"`

	// Type is the entity type for the Nintendo product.
	//
	// Type is always type game for Nintendo game products.
	Type string `json:"type"`
}

// compatabilityItem is the unfettered auxiliary information related to a Nintendo Amiibo product provided by nintendo.com.
//
// compatabilityItem describes the additional miscellaneous information that relates to Nintendo games that supports a Nintendo Amiibo product.
//
// compatabilityItem contains varying levels of completeness relative to the release status of Nintendo Amiibo products or game titles.
type compatabilityItem struct {

	// compatabilityCommon is a composed property.
	compatabilityCommon

	// Description is the description for the Nintendo product.
	//
	// Description is often a null string.
	Description string `json:"description"`

	// LastModified is the timestamp in milliseconds.
	LastModified int64 `json:"lastModified"`

	// Title is the name given to the Nintendo game item.
	//
	// Title often can contain special characters that will need filtering to prevent poor hashing keys.
	Title string `json:"title"`
}
