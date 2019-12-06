// Package amiibo handles creating, writing and organizing amiibo datasets into standardized formats.
//
// To create a standardized Amiibo (a standardized Amiibo is an Amiibo that is built of all
// of the common amiibo SDK components), the appropriate compatability and lineup
// information need to be found across all of the exposed fields in the various
// amiibo SDK XHR structs.
//
// This package does not assume the responsibility of performing this action (the data union mentioned above),
// and allows the developer to build their own custom union, passing in
// the relevant structs that help detail the level of verbosity each standardized Amiibo should export.
//
// If you want to defer to the amiibo SDKs implementation of this action,
// the amiibo/mix in conjunction with the amiibo/unmix can be used.
package amiibo
