package amiibo

type RawSlice []*RawAmiibo

func (pointer *RawSlice) Len() int {
	return len(*pointer)
}
