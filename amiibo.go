package main

type Amiibo struct {
	API       string  `json:"api"`
	Character string  `json:"character"`
	Game      string  `json:"game"`
	Head      string  `json:"head"`
	Image     string  `json:"image"`
	Name      string  `json:"name"`
	Release   Release `json:"release"`
	Series    string  `json:"series"`
	Type      string  `json:"type"`
}
