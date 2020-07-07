package amiibo

import "fmt"

type ENGGame struct{}

func (e ENGGame) AddENGChartGame(v ENGChartGame) (err error) {
	return
}

func (e ENGGame) AddENGChartItem(v ENGChartItem) (err error) {
	return
}

func NewENGGame(ENGChartGame ENGChartGame, ENGChartItem ENGChartItem) (v ENGGame, err error) {
	var ok bool
	ok = ENGChartGame.GetID() == ENGChartItem.GetID()
	if !ok {
		err = fmt.Errorf("ENGChartGame != ENGChartItem")
	}
	if err != nil {
		return
	}
	err = v.AddENGChartGame(ENGChartGame)
	if err != nil {
		return
	}
	err = v.AddENGChartItem(ENGChartItem)
	if err != nil {
		return
	}
	return
}
