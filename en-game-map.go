package amiibo

// ENGGameMap is map of ENGGame.
type ENGGameMap (map[string]ENGGame)

// NewENGGameMap returns a ENGGameMap.
func NewENGGameMap(ENGChart *ENGChart) (v ENGGameMap, err error) {
	v = (make(ENGGameMap))
	for _, EN := range ENGChart.GameList {
		ID := EN.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = ENGGame{}
		}
		var p = v[ID]
		err = (&p).AddENGChartGame(&EN)
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
		err = (&p).AddENGChartItem(&EN)
		if err != nil {
			return
		}
		v[ID] = p
	}
	return
}

// ReadENGGameMap reads a ENGGameMap from disc.
func ReadENGGameMap(dir string, filename string) (v ENGGameMap, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteENGGameMap writes a ENGGameMap to disc.
func WriteENGGameMap(dir string, filename string, v *ENGGameMap) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
