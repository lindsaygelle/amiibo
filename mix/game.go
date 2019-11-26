package mix

import (
	"fmt"

	"github.com/gellel/amiibo/compatability"
)

// Game is an aggregation of all game related data points across the
// various amiibo packages (amiibo/compatability, amiibo/lineup).
type Game struct {
	Game *compatability.Game
	Item *compatability.Item
}

func NewGame(g *compatability.Game, i *compatability.Item) (*Game, error) {
	var (
		game *Game
	)
	if g == nil {
		return nil, errGNil
	}
	if i == nil {
		return nil, errINil
	}
	if g.Key() != i.Key() {
		return nil, fmt.Errorf("*g and *i do not relate")
	}
	game = &Game{
		Game: g,
		Item: i}
	return game, nil
}
