package amiibo

// https://www.nintendo.co.jp/data/software/xml-system/amiibo-lineup-coming.xml
// https://www.nintendo.co.jp/hardware/amiibo/chart/data/lineup.xml
// https://www.nintendo.co.jp/hardware/amiibo/chart/data/chart.xml

// https://www.nintendo.com/content/noa/en_US/amiibo/line-up/jcr:content/root/responsivegrid/lineup.model.json
// https://www.nintendo.com/content/noa/en_US/amiibo/compatibility/jcr:content/root/responsivegrid/compatibility_chart.model.json

// amiibo is the normalized amiibo data scraped from a compatabilityAmiibo.
type amiibo struct{}

// amiiboChart is the unfettered Nintendo Amiibo chart information provided by nintendo.co.jp.
//
// amiiboChart contains the Japanese language Nintendo Amiibo software compatability.
//
// amiiboChart is assumed to be in Japanese hiragana.
type amiiboChart struct {
	Items []amiiboItem `xml:"items"`
}

// amiiboCompatibility is the unfettered Nintendo Amiibo compatibility information provided by nintendo.com.
//
// amiiboCompatibility contains the relationship information between Nintendo Amiibo products
// and the games or applications that can be used within.
//
// amiiboCompatibility is assumed to be in English.
type amiiboCompatibility struct {

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

// amiiboContent are the common properties shared between amiiboCompatibility and amiiboLineup.
type amiiboContent struct {

	// ComponentPath is the relative path to the Nintendo resource file.
	ComponentPath string `json:"componentPath"`
}

// amiiboItem is the unfettered Nintendo Amiibo chart information provided by nintendo.co.jp.
//
// amiiboItem contains the simplified Nintendo Amiibo product information.
//
// amiiboItem is assumed to be in Japanese Hiragana.
//
// amiiboItem is provided as XML from nintendo.co.jp.
type amiiboItem struct {
	chartCommon

	// Series is the Japanese Hiragana for the Nintendo product that the Nintendo Amiibo product is affiliated with.
	//
	// Series will need to be translated from Japanese to English.
	Series string `xml:"series"`

	// Softwares is a collection of metadata that the Nintendo Amiibo product integrates with.
	Softwares []amiiboItemSoftware `xml:"softwares"`
}

// amiiboItemSoftware is the software support information for a Nintendo Amiibo chart item provided by nintendo.co.jp.
//
// amiiboItemSoftware is assumed to be in Japanese Hiragana.
//
// amiiboItemSoftware is provided as XML from nintedo.co.jp.
type amiiboItemSoftware struct {

	// chartCommon is a composed property.
	chartCommon

	// More is the verbose description for the Nintendo Amiibo chart item.
	More string `xml:"more"`

	// Pickup is a provided property with an unclear purpose.
	Pickup int64 `xml:"pickup"`

	// ReadWrite is a provided property with an unclear purpose.
	ReadWrite string `xml:"readwrite"`
}

// amiiboLineup is the unfettered Nintendo Amiibo lineup information provided by nintendo.com.
//
// amiiboLineup contains the product information for the Nintendo Amiibo product.
//
// amiiboLineup is assumed to be in English.
type amiiboLineup struct {

	// amiiboContent is a composed property.
	amiiboContent

	// AmiiboList is a collection of Nintendo Amiibo products containing their product information.
	AmiiboList []lineupAmiibo

	// Items is a collection of metadata related to Nintendo Amiibo products.
	Items []lineupItem
}

type chartCommon struct {

	// Code is the English ID for the Nintendo Amiibo product from the Japanese CDN.
	Code string `xml:"code"`

	// Name is the name for the Nintendo product.
	//
	// Name is assumed to be in Japanese Hiragana but may be in English.
	Name string `xml:"name"`
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

type lineupAmiibo struct{}

type lineupItem struct{}
