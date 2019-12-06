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
// If you want to defer to the game SDKs implementation of this action,
// the amiibo SDK mix package in conjunction with the amiibo SDK unmix package can be used.
package game
