package compatability

// Amiibo is a representation of a Nintendo Amiibo product provided from resource.Compatability.
// Amiibo contains data provided as-is from Nintendo with a mixture of content
// provided for each Nintendo Amiibo product to describe its unique attributes.
// Amiibo provided from compatability.Amiibo focus on describing the
// figurines relational data with other Nintendo products. Amiibos
// collected from the compatability resource are consumed by the amiibo/amiibo
// package to construct a normalized aggregation of an Amiibo across all resources.
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
