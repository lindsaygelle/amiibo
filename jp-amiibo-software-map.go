package amiibo

// JPNAmiiboSoftwareMap is a map of JPNAmiiboSoftware.
type JPNAmiiboSoftwareMap (map[string]JPNAmiiboSoftware)

// Add adds a JPNAmiiboSoftware to the JPNAmiiboSoftwareMap.
func (j *JPNAmiiboSoftwareMap) Add(v *JPNAmiiboSoftware) (ok bool) {
	(*j)[v.GetID()] = *v
	ok = j.Has(v.GetID())
	return
}

// Del deletes a JPNAmiiboSoftware from the JPNAmiiboSoftwareMap.
func (j *JPNAmiiboSoftwareMap) Del(ID string) (ok bool) {
	delete(*j, ID)
	ok = j.Has(ID) == false
	return
}

// Each performs a for-each loop through the ENGAmiiboMap.
func (j *JPNAmiiboSoftwareMap) Each(f func(string, JPNAmiiboSoftware)) {
	for k, v := range *j {
		f(k, v)
	}
}

// Get gets an JPNAmiiboSoftware from the JPNAmiiboSoftwareMap.
func (j *JPNAmiiboSoftwareMap) Get(ID string) (v JPNAmiiboSoftware, ok bool) {
	v, ok = (*j)[ID]
	return
}

// Has checks if the JPNAmiiboSoftwareMap has a JPNAmiiboSoftware with the corresponding ID.
func (j *JPNAmiiboSoftwareMap) Has(ID string) (ok bool) {
	_, ok = j.Get(ID)
	return
}
