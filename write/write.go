package write

import (
	"os"

	"github.com/gellel/amiibo/amiibo"
	"github.com/gellel/amiibo/game"
)

const (
	// Permission is the os.FileMode all content is written as using the amiibo package.
	Permission os.FileMode = 0777
)

func Amiibo(amiibo *amiibo.Amiibo) {}

func Game(game *game.Game) {}
