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
