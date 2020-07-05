package amiibo

// JPNAmiiboMap is a map of JPNAmiibo.
type JPNAmiiboMap map[string]JPNAmiibo

func NewJPNAmiiboMap(JPNChart JPNChart, JPNLineup JPNLineup) (v JPNAmiiboMap) {
	go func() {
		for _, JP := range JPNChart.Items {
			ID := JP.GetID()
			if _, ok := v[ID]; ok {
				v[ID] = JPNAmiibo{}
			}
		}
	}()
	go func() {
		for _, JP := range JPNLineup.Items {
			ID := JP.GetID()
			if _, ok := v[ID]; ok {
				v[ID] = JPNAmiibo{}
			}
		}
	}()
	return
}
