package amiibo

import "fmt"

var (
	_ rawAmiiboPresentedBy = (*RawAmiiboPresentedBy)(nil)
)

type rawAmiiboPresentedBy interface{}

// A RawAmiiboPresentedBy string represents the publisher of the Amiibo figure.
type RawAmiiboPresentedBy string

func (r *RawAmiiboPresentedBy) String() string {
	return fmt.Sprintf("%s", *r)
}
