package amiibo

import "fmt"

// JPNSoftwareMap is a map of JPNSoftware.
type JPNSoftwareMap (map[string]JPNSoftware)

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
			err = fmt.Errorf("JPNSoftwareMap[JPNChartSoftware.Items.GetID()] != ok")
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
