package compatability

// Amiibo is a representation of a Nintendo Amiibo product provided from https://www.nintendo.com/amiibo/compatability/.
// Amiibo contains data provided as-is from Nintendo with a mixture of content
// provided for each Nintendo Amiibo product. Amiibo provided from compatability.Amiibo describes the
// Amiibo figure relational data with other Nintendo products.
type Amiibo struct {
	ID              string `json:"id"`
	Image           string `json:"image"`
	IsRelatedTo     string `json:"isRelatedTo"`
	IsReleased      string `json:"isReleased"`
	Name            string `json:"name"`
	ReleaseDateMask string `json:"releaseDateMask"`
	TagID           string `json:"tagid"`
	Type            string `json:"type"`
	URL             string `json:"url"`
}
