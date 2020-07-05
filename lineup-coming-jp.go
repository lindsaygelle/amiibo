package amiibo

// https://www.nintendo.co.jp/data/software/xml-system/amiibo-lineup-coming.xml

// lineupComingJPN is the unfettered upcoming Nintendo Amiibo product information provided by nintendo.co.jp.
type lineupComingJPN struct{}

// lineupComingItemJPN is the unfettered upcoming Nintendo Amiibo product information provided by nintendo.co.jp.
type lineupComingItemJPN struct {
	D              string `xml:"d"`
	Title          string `xml:"title"`
	ReleaseDateStr string `xml:"release_date_str"`
}
