package mix

import (
	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/errors"
)

// Game is an aggregation of all game related data points across the
// various amiibo packages (amiibo/compatability, amiibo/lineup).
type Game struct {
	Game *compatability.Game
	Item *compatability.Item
}

// NewGame creates a new instance of the mix.Game from the aggregation
// of amiibo structs across the amiibo package. Returns an error if all data points are
// not provided to the function or if a common union cannot be guaranteed across
// compatability.Game and compatability.Item.
func NewGame(g *compatability.Game, i *compatability.Item) (*Game, error) {
	var (
		game *Game
	)
	if g == nil {
		return nil, errors.ErrArgGNil
	}
	if i == nil {
		return nil, errors.ErrArgINil
	}
	if g.Key() != i.Key() {
		return nil, errors.ErrArgsNoRel
	}
	game = &Game{
		Game: g,
		Item: i}
	return game, nil
}
