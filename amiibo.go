package amiibo

// https://www.nintendo.com/content/noa/en_US/amiibo/compatibility/jcr:content/root/responsivegrid/compatibility_chart.model.json

// compatability is the unfettered Nintendo Amiibo compatibility information provided by nintendo.com.
//
// compatability contains the relationship information between Nintendo Amiibo products
// and the games or applications that can be used within.
//
// compatability is assumed to be in English.
type compatability struct {

	// amiiboContent is a composed property.
	amiiboContent

	// AuthorMode is the state of the Nintendo CDN for the compatability file.
	AuthorMode bool `json:"authorMode"`

	// AmiiboList is a collection of Nintendo Amiibo products that are compatible with Nintendo games.
	AmiiboList []compatabilityAmiibo `json:"amiiboList"`

	// Date is the date signature when the Nintendo compatability file was recevied.
	Date string `json:"date"`

	// ETag is the cache ID for the Nintendo compatability file.
	ETag string `json:"etag"`

	// GameList is a collection of Nintendo games that are compatible with Nintendo Amiibo products.
	GameList []compatabilityGame `json:"gameList"`

	// Items is a collection of related metadata that corresponds to Nintendo games that support Nintendo Amiibo.
	Items []compatabilityItem `json:"items"`

	// Language is the ISO key for the language the Nintendo information is presented in.
	Language string `json:"language"`

	// Mode is an exposed field from the Nintendo CDN that has an unknown purpose.
	Mode string `json:"mode"`
}

// amiiboContent are the common properties shared between compatability and lineup.
type amiiboContent struct {

	// ComponentPath is the relative path to the Nintendo resource file.
	ComponentPath string `json:"componentPath"`
}

// compatabilityAmiibo is the unfettered Nintendo Amiibo product data provided by nintendo.com.
//
// compatabilityAmiibo describes the abbreviated compatibility information for a specific Nintendo Amiibo figurine or card.
//
// compatabilityAmiibo contains varying levels of completeness relative to the release status of the product.
type compatabilityAmiibo struct {

	// compatabilityCommon is a composed property.
	compatabilityCommon
	// compatabilityCommonProduct is a composed property.
	compatabilityCommonProduct

	// IsRelatedTo is the name of a Nintendo product or series that the Nintendo Amiibo figurine or card is affiliated with.
	//
	// IsRelatedTo often can contain special characters that will need filtering to prevent poor hashing keys.
	IsRelatedTo string `json:"isRelatedTo"`

	// TagID is the identifier namespace for the Nintendo Amiibo figurine or card.
	TagID string `json:"tagId"`
}

// compatabilityCommon are the common properties shared between compatabilityAmiibo, compatabilityGame and compatabilityItem structs.
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

// compatabilityCommonProduct are the common product properties shared between compatabilityAmiibo and compatabilityGame.
type compatabilityCommonProduct struct {

	// ID is the unique identifier for the Nintendo product.
	//
	// ID is a UUID.
	ID string `json:"id"`

	// Image is the relative URL to the Nintendo game box art image.
	//
	// Image requires the nintendo.com domain prepended to the URL.
	Image string `json:"image"`

	// IsReleased is a string representation of a boolean that indicates whether the Nintendo product is available.
	//
	// IsReleased needs to be formatted to a bool data type.
	IsReleased string `json:"isReleased"`

	// Name is the name given to the Nintendo product.
	//
	// Name often can contain special characters that will need filtering to prevent poor hashing keys.
	Name string `json:"name"`

	// Type is the entity type for the Nintendo product.
	Type string `json:"type"`
}

// compatabilityGame is the unfettered game information related to a Nintendo Amiibo product provided by nintendo.com.
//
// compatabilityGame describes the abbreviated game product information that has known Nintendo Amiibo support.
//
// compatabilityGame contains varying levels of accuracy relative to the release status of Nintendo Amiibo products.
type compatabilityGame struct {

	// compatabilityCommon is a composed property.
	compatabilityCommon
	// compatabilityCommonProduct is a composed property.
	compatabilityCommonProduct
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
