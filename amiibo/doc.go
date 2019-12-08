// Package amiibo handles creating, writing and organizing amiibo datasets into standardized formats.
//
// To create a standardized Amiibo (a standardized Amiibo is an Amiibo that is composed of all
// of the common amiibo SDK components), the appropriate compatability and lineup
// information need to be found across all of the exposed fields in the various
// amiibo SDK XHR structs and merged into the one Amiibo struct.
//
// This package does not assume the responsibility of performing this action (the data composition mentioned above),
// and allows the developer to build their own custom union, passing in
// the relevant structs that help control the level of verbosity each 
// standardized Amiibo should export.
//
// The amiibo SDK provides a basic composition operation that will automatically search
// and parse the structs into a normalized Amiibo. Refer to the package documentation for
// more information.
package amiibo
