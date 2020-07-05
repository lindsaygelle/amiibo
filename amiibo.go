package amiibo

// https://www.nintendo.co.jp/data/software/xml-system/amiibo-lineup-coming.xml
// https://www.nintendo.co.jp/hardware/amiibo/chart/data/lineup.xml
// https://www.nintendo.co.jp/hardware/amiibo/chart/data/chart.xml

// https://www.nintendo.com/content/noa/en_US/amiibo/line-up/jcr:content/root/responsivegrid/lineup.model.json
// https://www.nintendo.com/content/noa/en_US/amiibo/compatibility/jcr:content/root/responsivegrid/compatibility_chart.model.json

// chart is the unfettered Nintendo Amiibo chart information provided by nintendo.co.jp.
//
// chart contains the Japanese language Nintendo Amiibo software compatability.
//
// chart is assumed to be in Japanese hiragana.
type chart struct {
	Items []chartAmiibo `xml:"items"`
}

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

// chartAmiibo is the unfettered Nintendo Amiibo chart information provided by nintendo.co.jp.
//
// chartAmiibo contains the simplified Nintendo Amiibo product information.
//
// chartAmiibo is assumed to be in Japanese Hiragana.
//
// chartAmiibo is provided as XML from nintendo.co.jp.
type chartAmiibo struct {

	// chartCommon is a composed property.
	chartCommon

	// Series is the Japanese Hiragana for the Nintendo product that the Nintendo Amiibo product is affiliated with.
	//
	// Series will need to be translated from Japanese to English.
	Series string `xml:"series"`

	// Softwares is a collection of metadata that the Nintendo Amiibo product integrates with.
	Softwares []chartSoftware `xml:"softwares"`
}

// chartSoftware is the software support information for a Nintendo Amiibo chart item provided by nintendo.co.jp.
//
// chartSoftware is assumed to be in Japanese Hiragana.
//
// chartSoftware is provided as XML from nintedo.co.jp.
type chartSoftware struct {

	// chartCommon is a composed property.
	chartCommon

	// More is the verbose description for the Nintendo Amiibo chart item.
	More string `xml:"more"`

	// Pickup is a provided property with an unclear purpose.
	Pickup int64 `xml:"pickup"`

	// ReadWrite is a provided property with an unclear purpose.
	ReadWrite string `xml:"readwrite"`
}

// chartCommon are the shared properties between chart
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

// lineup is the unfettered Nintendo Amiibo lineup information provided by nintendo.com.
//
// lineup contains the product information for the Nintendo Amiibo product.
//
// lineup is assumed to be in English.
type lineup struct {

	// amiiboContent is a composed property.
	amiiboContent

	// AmiiboList is a collection of Nintendo Amiibo products containing their product information.
	AmiiboList []lineupAmiibo

	// Items is a collection of metadata related to Nintendo Amiibo products.
	Items []lineupItem
}

// lineupJP is the unfettered Nintendo Amiibo lineup information provided by nintendo.co.jp.
//
// lineupJP contains the product properties related to Nintendo Amiibo products.
//
// lineupJP is assumed to be in Japanese Hiragana.
type lineupJP struct {

	// Item is a collection of Nintendo Amiibo products containing their product information in Japanese.
	Item []lineupItemJP `xml:"item"`

	// SeriesItem is a collection of Nintendo Amiibo product auxiliary information.
	SeriesItem []lineupSeriesItemJP `xml:"series_item"`
}

type lineupAmiibo struct{}

type lineupItem struct{}

type lineupItemJP struct {
	// BigSize is a integer representation of a boolean.
	//
	// BigSize relates to the scale of the Nintendo Amiibo product.
	BigSize int `xml:"bigsize"`

	// Chart is a integer representation of a boolean.
	//
	// Chart relates to the occurrence of the Nintendo Amiibo product in the chart XML.
	Chart int64 `xml:"chart"`

	// Code is the English ID for the Nintendo Amiibo product from the Japanese CDN.
	Code string `xml:"code"`

	// Date is the YYYYMMDD expression of the Nintendo Amiibo product release date.
	Date string `xml:"date"`

	// DisplayDate is the Japanese Hiragana expression of the Nintedo Amiibo product release date.
	//
	// DisplayDate currently (5/07/2020 (DD-MM-YYYY)) has a typo on the nintendo.co.jp and exists
	// as dispalydate.
	DisplayDate string `xml:"displayDate"`

	// Limited is a integer representation of a boolean.
	//
	// Limited relates to the rareness of the Nintendo Amiibo product.
	Limited int `xml:"limited"`

	// Name is the name of the Nintendo Amiibo product.
	//
	// Name is assumed to be in Japanese Hiragana but may be in English.
	Name string `xml:"name"`

	// NameKana is the name of the Nintendo Amiibo product in Japanese Hiragana.
	NameKana string `xml:"nameKana"`

	// New is a integer representation of a boolean.
	//
	// New relates to the newness of the Nintendo Amiibo product.
	New int `xml:"new"`

	// Price is the price of the Nintendo Amiibo product in Japanese Yen.
	Price string `xml:"price"`

	// Priority is the numerical rank of the Nintendo Amiibo product.
	Priority int64 `xml:"priority"`

	// Series is the Japanese Hiragana for the Nintendo product that the Nintendo Amiibo product is affiliated with.
	//
	// Series will need to be translated from Japanese to English.
	Series string `xml:"series"`
}

// lineupSeriesItemJP is the unfettered Nintendo Amiibo auxiliary metadata provided by nintendo.co.jp.
//
// lineupSeriesItemJP is assumed to be in Japanese Hiragana.
type lineupSeriesItemJP struct {

	// BGColor is the hexidecimal code for the Nintendo Amiibo product.
	BGColor string `xml:"bgcolor"`

	// Color is the hexidecimal code for the Nintendo Amiibo product.
	Color string `xml:"color"`

	// Name is the name of the Nintendo Amiibo product.
	//
	// Name is assumed to be in Japanese Hiragana but may be in English.
	Name string `xml:"name"`
}
