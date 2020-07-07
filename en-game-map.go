package amiibo

// ENGGGameMap is map of ENGGame.
type ENGGGameMap (map[string]ENGGame)

// NewENGGameMap returns a ENGGGameMap.
func NewENGGameMap(ENGChart ENGChart) (v ENGGGameMap, err error) {
	v = (make(ENGGGameMap))
	for _, EN := range ENGChart.GameList {
		ID := EN.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = ENGGame{}
		}
		var p = v[ID]
		err = (&p).AddENGChartGame(EN)
		if err != nil {
			return
		}
		v[ID] = p
	}
	for _, EN := range ENGChart.Items {
		ID := EN.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = ENGGame{}
		}
		var p = v[ID]
		err = (&p).AddENGChartItem(EN)
		if err != nil {
			return
		}
		v[ID] = p
	}
	return
}
