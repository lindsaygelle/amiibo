package amiibo

import "fmt"

// JPNAmiiboMap is a map of JPNAmiibo.
type JPNAmiiboMap (map[string]JPNAmiibo)

// Add adds a JPNAmiibo to the JPNAmiiboMap.
func (j *JPNAmiiboMap) Add(v *JPNAmiibo) (ok bool) {
	(*j)[v.GetID()] = *v
	ok = j.Has(v.GetID())
	return
}

// Del deletes a JPNAmiibo from the JPNAmiiboMap.
func (j *JPNAmiiboMap) Del(ID string) (ok bool) {
	delete(*j, ID)
	ok = j.Has(ID) == false
	return
}

// Each performs a for-each loop through the ENGAmiiboMap.
func (j *JPNAmiiboMap) Each(f func(string, JPNAmiibo)) {
	for k, v := range *j {
		f(k, v)
	}
}

// Get gets an JPNAmiibo from the JPNAmiiboMap.
func (j *JPNAmiiboMap) Get(ID string) (v JPNAmiibo, ok bool) {
	v, ok = (*j)[ID]
	return
}

// Has checks if the JPNAmiiboMap has a JPNAmiibo with the corresponding ID.
func (j *JPNAmiiboMap) Has(ID string) (ok bool) {
	_, ok = j.Get(ID)
	return
}

// Len returns the length of the JPNAmiiboMap.
func (j *JPNAmiiboMap) Len() int { return len(*j) }

// NewJPNAmiiboMap returns a new JPNAmiiboMap.
func NewJPNAmiiboMap(JPNChart *JPNChart, JPNLineup *JPNLineup) (v JPNAmiiboMap, err error) {
	v = (make(JPNAmiiboMap))
	for _, JP := range JPNChart.Items {
		ID := JP.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = JPNAmiibo{}
		}
		var p, ok = v[ID]
		if !ok {
			err = fmt.Errorf("(JPNAmiiboMap)[ID]: false")
		}
		if err != nil {
			return
		}
		err = (&p).AddJPNChartItem(&JP)
		if err != nil {
			return
		}
		v[ID] = p
	}
	for _, JP := range JPNLineup.Items {
		ID := JP.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = JPNAmiibo{}
		}
		var p, ok = v[ID]
		if !ok {
			err = fmt.Errorf("JPNAmiiboMap[JPNLineup.Items.GetID()] != ok")
		}
		if err != nil {
			return
		}
		err = (&p).AddJPNLineupItem(&JP)
		if err != nil {
			return
		}
		v[ID] = p
	}
	return
}

// ReadJPNAmiiboMap reads a JPNAmiiboMap from disc.
func ReadJPNAmiiboMap(dir string, filename string) (v JPNAmiiboMap, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteJPNAmiiboMap writes a JPNAmiiboMap to disc.
func WriteJPNAmiiboMap(dir string, filename string, v *JPNAmiiboMap) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, &v)
	return
}
