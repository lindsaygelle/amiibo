// Package amiibo is an unofficial SDK for the Go programming language.
//
// The amiibo SDK provides tools that developers can use to explore the Amiibo products
// on offer from Nintendo.
//
// The SDK is designed to remove the complexity of coding directly against the Nintendo Amiibo website
// and resolving the unification of data collected from the Nintendo CDN.
//
// The SDK also includes helpful utilities on top of the core struct layers that add additional
// capabilities and functionality. For example, the amiibo/mux package provides a
// basic HTTP mux interface that can be extended and hosted on an public service.
//
// To help break the package into smaller consumable and maintainable parts, most of the functionality
// has been separated into various subpackages. Each subpackage is intended to handle a particular
// level of complexity. To consume the amiibo directly from the Nintendo CDN, the amiibo/compatability
// and amiibo/lineup packages are the best places to start, as these items focus on contacting the
// Nintendo Amiibo website directly (through the use of the resources contained in the resources
// package).
//
// To consume the normalized Nintendo data, the amiibo/web package can be imported and used. The web
// package focuses on abstracting the task of fetching, merging and organizing all of
// the Nintendo data across all of the different resources.
package amiibo
