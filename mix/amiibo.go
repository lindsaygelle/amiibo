package mix

import (
	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/errors"
	"github.com/gellel/amiibo/lineup"
)

// Amiibo is an aggregation of all Amiibo related data points across the
// various amiibo packages (amiibo/compatability, amiibo/lineup).
type Amiibo struct {
	Compatability *compatability.Amiibo
	Item          *lineup.Item
	Lineup        *lineup.Amiibo
}

// NewAmiibo creates a new instance of the mix.Amiibo from the aggregation
// of amiibo structs across the amiibo package. Returns an error if all data points are
// not provided to the function or if a common union cannot be guaranteed across
// compatability.Amiibo, lineup.Item and lineup.Amiibo.
func NewAmiibo(c *compatability.Amiibo, i *lineup.Item, l *lineup.Amiibo) (*Amiibo, error) {
	var (
		amiibo *Amiibo
	)
	if c == nil {
		return nil, errors.ErrArgCNil
	}
	if i == nil {
		return nil, errors.ErrArgINil
	}
	if l == nil {
		return nil, errors.ErrArgLNil
	}
	if c.Key() != i.Key() || i.Key() != l.Key() {
		return nil, errors.ErrArgsNoRel
	}
	amiibo = &Amiibo{
		Compatability: c,
		Item:          i,
		Lineup:        l}
	return amiibo, nil
}
