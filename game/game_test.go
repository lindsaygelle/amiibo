package game_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/gellel/amiibo/game"
	"github.com/gellel/amiibo/mix"
)

func Test(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	m, err := mix.Get()
	if err != nil {
		panic(err)
	}
	g := game.NewFromMix(m.Games)
	b, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}
	filename := fmt.Sprintf("%s.json", filepath.Join(u.HomeDir, "Desktop", "Game"))
	err = ioutil.WriteFile(filename, b, 0777)
	if err != nil {
		panic(err)
	}
}
