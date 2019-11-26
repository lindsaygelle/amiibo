package compatability

// Game is a representation of a Nintendo video-game product provided from resource.Compatability.
// Game contains data provided as-is from Nintendo with a mixture of content
// that provides a reference for a Nintendo video-game with a Nintendo Amiibo product.
// Games collected from the compatability resource are consumed by the amiibo/game
// package to construct a normalized aggregation of a Nintendo video-game across all resources.
type Game struct {
	Image           string `json:"image"`
	ID              string `json:"id"`
	IsReleased      string `json:"isReleased"`
	Name            string `json:"name"`
	Path            string `json:"path"`
	ReleaseDateMask string `json:"releaseDateMask"`
	Type            string `json:"type"`
	URL             string `json:"url"`
}
