package mix

import (
	"fmt"

	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/lineup"
)

// Amiibo is an aggregation of all Amiibo related data points across the
// various amiibo packages (amiibo/compatability, amiibo/lineup).
type Amiibo struct {
	Compatability *compatability.Amiibo
	Item          *lineup.Item
	Lineup        *lineup.Amiibo
}

func NewAmiibo(c *compatability.Amiibo, i *lineup.Item, l *lineup.Amiibo) (*Amiibo, error) {
	var (
		amiibo *Amiibo
	)
	if c == nil {
		return nil, errCNil
	}
	if i == nil {
		return nil, errINil
	}
	if l == nil {
		return nil, errLNil
	}
	if c.Key() != i.Key() || i.Key() != l.Key() {
		return nil, fmt.Errorf("*c, *i and *l do not relate")
	}
	amiibo = &Amiibo{
		Compatability: c,
		Item:          i,
		Lineup:        l}
	return amiibo, nil
}
