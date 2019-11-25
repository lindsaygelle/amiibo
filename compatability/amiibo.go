package compatability

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
