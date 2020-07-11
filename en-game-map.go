package amiibo

import "fmt"

// ENGGameMap is map of ENGGame.
type ENGGameMap (map[string]ENGGame)

// Add adds a ENGGame to the ENGGameMap.
func (e *ENGGameMap) Add(v *ENGGame) (ok bool) {
	(*e)[v.GetID()] = *v
	ok = e.Has(v.GetID())
	return
}

// Del deletes a ENGGame from the ENGGameMap.
func (e *ENGGameMap) Del(ID string) (ok bool) {
	delete(*e, ID)
	ok = e.Has(ID) == false
	return
}

// Each performs a for-each loop through the ENGGameMap.
func (e *ENGGameMap) Each(f func(string, ENGGame)) {
	for k, v := range *e {
		f(k, v)
	}
}

// Get gets an ENGGame from the ENGGameMap.
func (e *ENGGameMap) Get(ID string) (v ENGGame, ok bool) {
	v, ok = (*e)[ID]
	return
}

// Has checks if the ENGGameMap has a ENGGame with the corresponding ID.
func (e *ENGGameMap) Has(ID string) (ok bool) {
	_, ok = e.Get(ID)
	return
}

// NewENGGameMap returns a ENGGameMap.
func NewENGGameMap(ENGChart *ENGChart) (v ENGGameMap, err error) {
	v = (make(ENGGameMap))
	for _, EN := range ENGChart.GameList {
		ID := EN.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = ENGGame{}
		}
		var p, ok = v[ID]
		if !ok {
			err = fmt.Errorf("(ENGGameMap)[ID]: false")
		}
		if err != nil {
			return
		}
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
