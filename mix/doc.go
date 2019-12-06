// Package mix is the amiibo SDK struct parser for organizing the Nintendo HTTP response content.
//
// This package handles the collection of the various Nintendo Amiibo data points into
// mixed content structures that can be later consumed by the amiibo SDK package.
//
// The provided mechanism assumes that each Nintendo HTTP response contains at least
// one reference to another structure that indicates a union between the two.
//
// This package is intended to help abstract the search process for each of the Nintendo HTTP response
// into a single function so that the focus can be on building the application around
// the mixed content structure.
package mix
