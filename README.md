![Go](https://github.com/lindsaygelle/amiibo/workflows/Go/badge.svg)
[![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/lindsaygelle/amiibo/blob/master/LICENSE)

# Amiibo

Amiibo is the an Nintendo Amiibo SDK for the Go programming language.

Amiibo is a fan project to collect the latest information about the Amiibo products on offer
from Nintendo. Package is built on scraping the Nintendo Amiibo website and normalizing
the data into consumable chunks.

The package is built around the Go API reference documentation. Please consider using `godoc`
to build custom integrations. If you are using Go 1.12 or earlier, godoc should be included. All
Go 1.13 users will need to grab this package using the `go get` flow.

## Installing

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or project's Go module dependencies.

```go get github.com/lindsaygelle/amiibo```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

```go get -u github.com/lindsaygelle/amiibo```

## Dependencies

The Amiibo SDK is built Go's module system. Please update to Go 1.13 to use `go mod` out of the box. 

## Go Modules

If you are using Go modules, your go get will default to the latest tagged release version of the SDK. To get a specific release version of the SDK use `@<tag>` in your `go get` command.

```go get github.com/lindsaygelle/amiibo@<version>```

To get the latest SDK repository change use @latest.

## License

This SDK is distributed under the Apache License, Version 2.0, see LICENSE.txt and NOTICE.txt for more information.