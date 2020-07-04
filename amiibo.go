package amiibo

// amiibo is the normalized amiibo data scraped from a rawAmiibo.
type amiibo struct{}

// amiiboCompatability is the unfettered Nintendo Amiibo compatability information provided by nintendo.com.
//
// amiiboCompatability contains the relationship information between Nintendo Amiibo products
// and the games or applications that can be used within.
type amiiboCompatability struct {
	AuthorMode    bool        `json:"authorMode"`
	AmiiboList    []rawAmiibo `json:"amiiboList"`
	ComponentPath string      `json:"componentPath"`
	GameList      []rawGame   `json:"gameList"`
	Items         []rawItem   `json:"items"`
	Language      string      `json:"language"`
	Mode          string      `json:"mode"`
}

// rawAmiibo is the unfettered amiibo product data provided by nintendo.com.
//
// rawAmiibo describes the abbreviated compatability information for a specific amiibo figurine or card.
type rawAmiibo struct{}

// rawGame is the unfettered game information related to a Nintendo amiibo product provided by nintendo.com.
//
// rawGame describes the abbreviated game product information that has Nintendo amiibo support.
type rawGame struct{}

