// Package dir is a basic directory handler for the amiibo package.
//
// Dir handles the creation, removal and checking of directories
// across the package in an effort to provide a single touchpoint
// for folder creation.
//
// Dir is consumed by amiibo/file to aid in the checking of
// path arguments to ensure that the intended destination is
// a directory as well as assist in the creation of utility
// folders to offer some out of the box functionality for the Amiibo
// content collected from Nintendo.
package dir
