package mix

import (
	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/lineup"
)

// Amiibo is an aggregation of all Amiibo related data points across the
// various amiibo packages (amiibo/compatability, amiibo/lineup).
type Amiibo struct {
	Compatability *compatability.Amiibo
	Item          *lineup.Item
	lineup        *lineup.Amiibo
}
