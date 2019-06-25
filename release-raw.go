package amiibo

func newRawRelease() *RawRelease {
	return &RawRelease{}
}

func NewRawRelease(AU, EU, JP, NA string) *RawRelease {
	return &RawRelease{
		AU: AU,
		EU: EU,
		JP: JP,
		NA: NA}
}

type RawRelease struct {
	AU string `json:"AU"`
	EU string `json:"EU"`
	JP string `json:"JP"`
	NA string `json:"NA"`
}
