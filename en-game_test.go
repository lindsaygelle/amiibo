package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engGameFileName = "eng-game-test.json"

func TestENGGame(t *testing.T) {
	var err error
	var ENGGame = amiibo.ENGGame{}
	_, err = amiibo.WriteENGGame(filefolder, engGameFileName, &ENGGame)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGGame(filefolder, engGameFileName)
	if err != nil {
		t.Fatal(err)
	}
}
