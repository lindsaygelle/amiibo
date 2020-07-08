package amiibo

// ENGAmiiboMap is map of ENGAmiibo.
type ENGAmiiboMap (map[string]ENGAmiibo)

// Add adds a ENGAmiibo to the ENGAmiiboMap.
func (e *ENGAmiiboMap) Add(v *ENGAmiibo) {
	(*e)[(*v).GetID()] = *v
}

// NewENGAmiiboMap returns a ENGAmiiboMap.
func NewENGAmiiboMap(ENGChart ENGChart, ENGLineup ENGLineup) (v ENGAmiiboMap, err error) {
	v = (make(ENGAmiiboMap))
	for _, EN := range ENGChart.AmiiboList {
		ID := EN.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = ENGAmiibo{}
		}
		var p = v[ID]
		err = (&p).AddENGChartAmiibo(&EN)
		if err != nil {
			return
		}
		v[ID] = p
	}
	for _, EN := range ENGLineup.AmiiboList {
		ID := EN.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = ENGAmiibo{}
		}
		var p = v[ID]
		err = (&p).AddENGLineupAmiibo(&EN)
		if err != nil {
			return
		}
		v[ID] = p
	}
	for _, EN := range ENGLineup.Items {
		ID := EN.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = ENGAmiibo{}
		}
		var p = v[ID]
		err = (&p).AddENGLineupItem(&EN)
		if err != nil {
			return
		}
		v[ID] = p
	}
	return
}

// ReadENGAmiiboMap reads a ENGAmiiboMap from disc.
func ReadENGAmiiboMap(dir string, filename string) (v ENGAmiiboMap, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteENGAmiiboMap writes a ENGAmiiboMap to disc.
func WriteENGAmiiboMap(dir string, filename string, v *ENGAmiiboMap) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
