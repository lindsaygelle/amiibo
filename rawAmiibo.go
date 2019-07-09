package amiibo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	_ rawAmiibo = (*RawAmiibo)(nil)
)

func deleteRawAmiibo(rawAmiibo *RawAmiibo) error {
	return os.Remove(filepath.Join(storepathRawAmiibo(), fmt.Sprintf("%s.json", rawAmiibo.HexCode)))
}

func getRawAmiibo(ID string) *RawAmiibo {
	if ok := strings.HasSuffix(ID, ".json"); !ok {
		ID = fmt.Sprintf("%s.json", ID)
	}
	if ok := strings.HasPrefix(ID, "r-"); !ok {
		ID = fmt.Sprintf("r-%s", ID)
	}
	b, err := openRawAmiibo(ID)
	if err != nil {
		return nil
	}
	amiibo, err := unmarshallRawAmiibo(b)
	if err != nil {
		return nil
	}
	return amiibo
}

func marshallRawAmiibo(rawAmiibo *RawAmiibo) ([]byte, error) {
	content, err := json.Marshal(rawAmiibo)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func openRawAmiibo(name string) (*[]byte, error) {
	filepath := filepath.Join(storepathRawAmiibo(), name)
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(reader)
	defer reader.Close()
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func storepathRawAmiibo() string {
	return filepath.Join(rootpath, "amiibo")
}

func unmarshallRawAmiibo(content *[]byte) (*RawAmiibo, error) {
	r := &RawAmiibo{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func writeRawAmiibo(rawAmiibo *RawAmiibo) error {
	err := os.MkdirAll(storepathRawAmiibo(), 0644)
	if err != nil {
		return err
	}
	content, err := marshallRawAmiibo(rawAmiibo)
	if err != nil {
		return err
	}
	filepath := filepath.Join(storepathRawAmiibo(), fmt.Sprintf("r-%s.json", rawAmiibo.HexCode))
	return ioutil.WriteFile(filepath, content, 0644)
}

func newRawAmiibo(r *json.RawMessage) *RawAmiibo {
	rawAmiibo := &RawAmiibo{}
	err := json.Unmarshal(*r, rawAmiibo)
	if err != nil {
		return nil
	}
	return rawAmiibo
}

type rawAmiibo interface {
	String() string
}

// A RawAmiibo struct contains the as-is provided data found in the scraped Nintendo Amiibo XHR HTTP response.
type RawAmiibo struct {
	AmiiboName          string `json:"amiiboName"`          // "Toon Link - The Wind Waker"
	AmiiboPage          string `json:"amiiboPage"`          // "/amiibo/detail/toon-link…30th-anniversary-series"
	BoxArtURL           string `json:"boxArtUrl"`           // "/content/dam/noa/en_US/a…nk-WW_30thAnniv_box.png"
	DetailsPath         string `json:"detailsPath"`         // "/content/noa/en_US/amiib…30th-anniversary-series"
	DetailsURL          string `json:"detailsUrl"`          // "/amiibo/detail/toon-link…30th-anniversary-series"
	FigureURL           string `json:"figureURL"`           // "/content/dam/noa/en_US/a…k-WW_30thAnniv_char.png"
	Franchise           string `json:"franchise"`           // "The Legend of Zelda"
	GameCode            string `json:"gameCode"`            // "NVLEAK2A"
	HexCode             string `json:"hexCode"`             // "#ffdc81"
	IsReleased          bool   `json:"isReleased"`          // true
	OverviewDescription string `json:"overviewDescription"` // "<p style=\"margin-top: 0…ol the winds.\n</p>"
	PresentedBy         string `json:"presentedBy"`         // "noa:publisher/nintendo"
	Price               string `json:"price"`               // "24.99"
	ReleaseDateMask     string `json:"releaseDateMask"`     // "12/02/2016"
	Series              string `json:"series"`              // "30th anniversary"
	Slug                string `json:"slug"`                // "toon-link-the-wind-waker…30th-anniversary-series"
	Type                string `json:"type"`                // "Figure"
	UnixTimestamp       int64  `json:"unixTimestamp"`       // 1480636800
	UPC                 string `json:"upc"`                 // "045496893064"
}

func (pointer *RawAmiibo) String() string {
	return fmt.Sprintf("%s", pointer.AmiiboName)
}
