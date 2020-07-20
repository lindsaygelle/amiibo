package amiibo

import "fmt"

// JPNSoftwareMap is a map of JPNSoftware.
type JPNSoftwareMap (map[string]JPNSoftware)

// Add adds a JPNSoftware to the JPNSoftwareMap.
func (j *JPNSoftwareMap) Add(v *JPNSoftware) (ok bool) {
	(*j)[v.GetID()] = *v
	ok = j.Has(v.GetID())
	return
}

// Del deletes a JPNSoftware from the JPNSoftwareMap.
func (j *JPNSoftwareMap) Del(ID string) (ok bool) {
	delete(*j, ID)
	ok = j.Has(ID) == false
	return
}

// Each performs a for-each loop through the JPNSoftwareMap.
func (j *JPNSoftwareMap) Each(f func(string, JPNSoftware)) {
	for k, v := range *j {
		f(k, v)
	}
}

// Get gets an JPNSoftware from the JPNSoftwareMap.
func (j *JPNSoftwareMap) Get(ID string) (v JPNSoftware, ok bool) {
	v, ok = (*j)[ID]
	return
}

// Has checks if the JPNSoftwareMap has a JPNSoftware with the corresponding ID.
func (j *JPNSoftwareMap) Has(ID string) (ok bool) {
	_, ok = j.Get(ID)
	return
}

// Len returns the length of the JPNSoftwareMap.
func (j *JPNSoftwareMap) Len() int { return len(*j) }

// NewJPNSoftwareMap returns a new JPNSoftwareMap.
func NewJPNSoftwareMap(JPNChartSoftware *JPNChartSoftware) (v JPNSoftwareMap, err error) {
	v = (make(JPNSoftwareMap))
	for _, JP := range JPNChartSoftware.Items {
		ID := JP.GetID()
		if _, ok := v[ID]; !ok {
			v[ID] = JPNSoftware{}
		}
		var p, ok = v[ID]
		if !ok {
			err = fmt.Errorf("(JPNChartSoftware)[ID]: false")
		}
		if err != nil {
			return
		}
		err = (&p).AddJPNChartSoftwareItem(&JP)
		if err != nil {
			return
		}
		v[ID] = p
	}
	return
}

// ReadJPNSoftwareMap reads a JPNSoftwareMap from disc.
func ReadJPNSoftwareMap(dir string, filename string) (v JPNSoftwareMap, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteJPNSoftwareMap writes a JPNSoftwareMap to disc.
func WriteJPNSoftwareMap(dir string, filename string, v *JPNSoftwareMap) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, &v)
	return
}
