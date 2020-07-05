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

	// DetailsPath is the relative path to the Nintendo Amiibo product information page.
	DetailsPath string `json:"detailsPath"`

	// DetailsURL is the relative URL to the Nintendo Amiibo product information page.
	DetailsURL string `json:"detailsUrl"`

	// UnixTimestamp is the Nintendo Amiibo product release date in milliseconds.
	UnixTimestamp int64 `json:"unixTimestamp"`
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
