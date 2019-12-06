// Package game handles creating, writing and organizing game datasets into standardized formats.
//
// To create a standardized Game (a standardized game is a snapshot of
// a Nintendo video game that is composed of all
// of the common game SDK components), the appropriate compatability and lineup
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
package game
