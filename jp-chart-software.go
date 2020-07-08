package amiibo

import "encoding/xml"

// https://www.nintendo.co.jp/hardware/amiibo/chart/data/software.xml

// JPNChartSoftware is the unfettered Japanese language Nintendo software information.
type JPNChartSoftware struct {
	XMLName xml.Name `xml:"items"`

	Items []JPNChartSoftwareItem `xml:"items>item"`
}
