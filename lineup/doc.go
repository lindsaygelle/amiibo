// Package lineup provides the basic package tools for contacting,
// collecting and organising the content from the Nintendo Amiibo
// lineup HTTP response.
//
// The package primarily exports methods on interfacing with the
// exported XHR struct, managing the normalization, writing
// and loading of the data from the CDN to the local OS.
//
// Lineup also exports a number of structures that help organise the content
// returned from the HTTP response.
//
// The Amiibo struct is provided to organise the unique traits and
// attributes of an individual Nintendo Amiibo figurine into a consumbale
// single touch data structure. These Amiibo structs
// are built on the fields provided by Nintendo and are likely to change
// as Nintendo maintains their content. Please raise any issues
// should this occur.
//
// Each Amiibo struct provided by Nintendo is currently generated to help
// in describing the Amiibo as an individual
// product, expressing the background meta information for each item.
//
// The use of the Amiibo struct can be used to build out any implementation of a data
// wrangler that makes the fields of the struct consumable by another program,
// but is important to keep in mind that the Amiibo (as of writing this
// documentation) does not contain all the product information and
// needs to be used in conjunction with the compatability package to build out the
// full picture of an Amiibo product.
package lineup
