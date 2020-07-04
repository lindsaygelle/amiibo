// Package amiibo is an unofficial SDK for the Go programming language.
package amiibo

// amiibo is the normalized amiibo data scraped from a rawAmiibo.
type amiibo struct{}

// amiiboCompatability is the unfettered Nintendo Amiibo compatability information provided by nintendo.com.
type amiiboCompatability struct {
	AmiiboList    []rawAmiibo `json:"amiiboList"`
	ComponentPath string      `json:"componentPath"`
	Language      string      `json:"language"`
}

// rawAmiibo is the unfettered amiibo product data provided by nintendo.com.
type rawAmiibo struct{}

