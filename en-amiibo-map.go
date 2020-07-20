package amiibo

import "fmt"

// ENGAmiiboMap is map of ENGAmiibo.
type ENGAmiiboMap (map[string]ENGAmiibo)

// Add adds a ENGAmiibo to the ENGAmiiboMap.
func (e *ENGAmiiboMap) Add(v *ENGAmiibo) (ok bool) {
	(*e)[v.GetID()] = *v
	ok = e.Has(v.GetID())
	return
}

// Del deletes a ENGAmiibo from the ENGAmiiboMap.
func (e *ENGAmiiboMap) Del(ID string) (ok bool) {
	delete(*e, ID)
	ok = e.Has(ID) == false
	return
}

// Each performs a for-each loop through the ENGAmiiboMap.
func (e *ENGAmiiboMap) Each(f func(string, ENGAmiibo)) {
	for k, v := range *e {
		f(k, v)
	}
}

// Get gets an ENGAmiibo from the ENGAmiiboMap.
func (e *ENGAmiiboMap) Get(ID string) (v ENGAmiibo, ok bool) {
	v, ok = (*e)[ID]
	return
}

// Has checks if the ENGAmiiboMap has a ENGAmiibo with the corresponding ID.
func (e *ENGAmiiboMap) Has(ID string) (ok bool) {
	_, ok = e.Get(ID)
	return
}

// Len returns the length of the ENGAmiiboMap.
func (e *ENGAmiiboMap) Len() int { return len(*e) }

// NewENGAmiiboMap returns a ENGAmiiboMap.
func NewENGAmiiboMap(ENGChart ENGChart, ENGLineup ENGLineup) (v ENGAmiiboMap, err error) {
	v = (make(ENGAmiiboMap))
	for _, EN := range ENGChart.AmiiboList {
		ID := EN.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = ENGAmiibo{}
		}
		var p, ok = v[ID]
		if !ok {
			err = fmt.Errorf("(ENGAmiiboMap)[ID]: false")
		}
		if err != nil {
			return
		}
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
		var p, ok = v[ID]
		if !ok {
			err = fmt.Errorf("(ENGAmiiboMap)[ID]: false")
		}
		if err != nil {
			return
		}
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
		var p, ok = v[ID]
		if !ok {
			err = fmt.Errorf("(ENGAmiiboMap)[ID]: false")
		}
		if err != nil {
			return
		}
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
