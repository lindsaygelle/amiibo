package amiibo

// JPNAmiiboSoftwareMap is a map of JPNAmiiboSoftware.
type JPNAmiiboSoftwareMap (map[string]JPNAmiiboSoftware)

// Add adds a JPNAmiiboSoftware to the JPNAmiiboSoftwareMap.
func (j *JPNAmiiboSoftwareMap) Add(v *JPNAmiiboSoftware) {
	(*j)[v.GetID()] = *v
}
