package amiibo

var (
	_ rawSlice = (*RawSlice)(nil)
)

type rawSlice interface {
	Len() int
}

// RawSlice is a collection of Raw Amiibo structs fetched from the Amiibo API endpoint.
type RawSlice []*RawAmiibo

func (pointer *RawSlice) Len() int {
	return len(*pointer)
}
