// Package amiibo handles creating, writing and organizing amiibo datasets into standardized formats.
//
// To create a standardized Amiibo (a standardized Amiibo is an Amiibo that is composed of all
// of the common amiibo SDK components), the appropriate compatability and lineup
// information need to be found across all of the exposed fields in the various
// amiibo SDK XHR structs.
//
// This package does not assume the responsibility of performing this action (the data union mentioned above),
// and allows the developer to build their own custom union, passing in
// the relevant structs that help detail the level of verbosity each standardized Amiibo should export.
//
// The amiibo SDK provides a basic union operation that will automatically compose
// and parse the structs into a normalized game. Refer to the package documenation for
// more information.
package amiibo
