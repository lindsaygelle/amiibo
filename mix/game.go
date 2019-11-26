package mix

import (
	"github.com/gellel/amiibo/compatability"
)

// Game is an aggregation of all game related data points across the
// various amiibo packages (amiibo/compatability, amiibo/lineup).
type Game struct {
	Item *compatability.Item
	Game *compatability.Game
}
