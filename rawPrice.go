package amiibo

import (
	"fmt"

	"golang.org/x/text/currency"
)

func NewAmiiboPrice(rawAmiiboPrice *RawAmiiboPrice) *currency.Amount {
	return new(currency.Amount)
}

var (
	_ rawAmiiboPrice = (*RawAmiiboPrice)(nil)
)

type rawAmiiboPrice interface{}

// A RawAmiiboPrice string represents the cost of an Amiibo (in USD) found in RawAmiibo within the Nintendo XHR HTTP response.
type RawAmiiboPrice string

func (r *RawAmiiboPrice) String() string {
	return fmt.Sprintf("%s", *r)
}
