package amiibo

import "fmt"

var (
	_ rawResponse = (*RawResponse)(nil)
)

type rawResponse interface {
	String() string
}

type RawResponse struct {
	Amiibo *RawSlice `json:"amiibo"`
}

func (pointer *RawResponse) String() string {
	return fmt.Sprintf("%v", pointer.Amiibo)
}
