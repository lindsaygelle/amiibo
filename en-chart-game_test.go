package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engChartGameFileName = "eng-chart-game-test.json"

func TestENGGameAmiibo(t *testing.T) {
	var err error
	var ENGGameAmiibo = amiibo.ENGChartGame{}
	_, err = amiibo.WriteENGChartGame(filefolder, engChartGameFileName, &ENGGameAmiibo)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGChartGame(filefolder, engChartGameFileName)
	if err != nil {
		t.Fatal(err)
	}
}
