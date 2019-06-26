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

type RawAmiiboPrice string

func (r *RawAmiiboPrice) String() string {
	return fmt.Sprintf("%s", *r)
}
