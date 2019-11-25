package mix

import (
	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/lineup"
)

type Amiibo struct {
	Compatability *compatability.Amiibo
	Item          *lineup.Item
	lineup        *lineup.Amiibo
}
