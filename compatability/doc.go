// Package compatability provides the basic package tools for contacting,
// collecting and organising the raw content output from the Nintendo Amiibo
// compatability content store.
//
// Compatability exports an Amiibo structure that describes the
// relational attributes of a Nintendo Amiibo figurine in the context
// of the Nintendo video games that are available for various platforms.
//
// The use of the Amiibo struct can be used to build out any implementation of a data
// wrangler that makes the fields of the struct consumable by another program,
// but is important to keep in mind that the Amiibo (as of writing this
// documentation) does not contain all the figurines product information and
// needs to be used in conjunction with the lineup package to built out the
// full picture.
package compatability
