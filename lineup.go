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
type lineupAmiibo struct{}

// lineupItem is the unfettered Nintendo Amiibo product additional information from nintendo.com.
//
// lineupItem contains additional information for a Nintendo Amiibo product.
type lineupItem struct {

	// Description is the verbose Nintendo Amiibo product summary.
	Description string `json:"description"`

	// LastModified is the Nintendo Amiibo product release date in milliseconds.
	LastModified int64 `json:"lastModified"`

	// Path is the relative path to the Nintendo Amiibo product information page according to the nintendo.com CDN.
	Path string `json:"path"`

	// Title is the name of the Nintendo Amiibo product.
	Title string `json:"title"`

	// URL is the relative path URL to the Nintendo Amiibo product information page.
	URL string `json:"url"`
}
