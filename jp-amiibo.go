package amiibo

import "fmt"

// JPNAmiibo is a formatted JPN Nintendo Amiibo.
type JPNAmiibo struct {

	// BigSize relates to the scale of the Nintendo Amiibo product.
	BigSize bool `json:"bigsize"`

	// Chart is a integer representation of a boolean.
	//
	// Chart relates to the occurrence of the Nintendo Amiibo product in the chart json.
	Chart int64 `json:"chart"`

	// Code is the English ID for the Nintendo Amiibo product from the Japanese CDN.
	Code string `json:"code"`

	// Date is the YYYYMMDD expression of the Nintendo Amiibo product release date.
	Date string `json:"date"`

	// DisplayDate is the Japanese Hiragana expression of the Nintedo Amiibo product release date.
	//
	// DisplayDate currently (5/07/2020 (DD-MM-YYYY)) has a typo on the nintendo.co.jp and exists
	// as dispalydate.
	DisplayDate string `json:"displayDate"`

	// Limited is a integer representation of a boolean.
	//
	// Limited relates to the rareness of the Nintendo Amiibo product.
	Limited bool `json:"limited"`

	Software []JPNAmiiboSoftware `json:"software"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `json:"name"`

	// NameKana is the name of the Nintendo Amiibo product in Japanese Hiragana.
	NameKana string `json:"nameKana"`

	// New is a integer representation of a boolean.
	//
	// New relates to the newness of the Nintendo Amiibo product.
	New bool `json:"new"`

	// Price is the price of the Nintendo Amiibo product in Japanese Yen.
	Price string `json:"price"`

	// Priority is the numerical rank of the Nintendo Amiibo product.
	Priority int64 `json:"priority"`

	// Series is the Japanese Hiragana for the Nintendo product that the Nintendo Amiibo product is affiliated with.
	Series string `json:"series"`
}

func newJPNAmiibo(a JPNChartItem, b JPNLineupItem) (v JPNAmiibo, err error) {
	if len(a.GetID()) == 0 {
		err = fmt.Errorf("JPNAmiibo: len(a.GetID())")
	}
	if err != nil {
		return
	}
	if len(b.GetID()) == 0 {
		err = fmt.Errorf("JPNAmiibo: len(b.GetID())")
	}
	if err != nil {
		return
	}
	if a.GetID() != b.GetID() {
		err = fmt.Errorf("JPNAmiibo: a.GetID() != b.GetID()")
	}
	if err != nil {
		return
	}
	return
}
