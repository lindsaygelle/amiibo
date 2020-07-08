package amiibo

// JPNAmiiboMap is a map of JPNAmiibo.
type JPNAmiiboMap (map[string]JPNAmiibo)

// NewJPNAmiiboMap returns a new JPNAmiiboMap.
func NewJPNAmiiboMap(JPNChart JPNChart, JPNLineup JPNLineup) (v JPNAmiiboMap, err error) {
	v = (make(JPNAmiiboMap))
	for _, JP := range JPNChart.Items {
		ID := JP.GetID()
		if _, ok := v[ID]; ok {
			v[ID] = JPNAmiibo{}
		}
		var p = v[ID]
		err = (&p).AddJPNChartItem(JP)
		if err != nil {
			return
		}
		v[ID] = p
	}
	for _, JP := range JPNLineup.Items {
		ID := JP.GetID()
		if _, ok := v[ID]; ok {
			v[ID] = JPNAmiibo{}
		}
		var p = v[ID]
		err = (&p).AddJPNLineupItem(JP)
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
func WriteJPNAmiiboMap(dir string, filename string, v JPNAmiiboMap) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, &v)
	return
}
