# findup.go [![Build Status](https://travis-ci.org/h2non/findup.png)](https://travis-ci.org/h2non/findup) [![GitHub release](https://img.shields.io/github/tag/h2non/findup.svg)]()

A [Go](http://golang.org) tiny package to **find the first file matching** in the current directory or the nearest ancestor directory up to root, with additional [Glob patterns](http://en.wikipedia.org/wiki/Glob_%28programming%29) support

It works in Go >= 1.1. See Travis [builds](https://travis-ci.org/h2non/findup)

## Installation

```bash
go get github.com/h2non/findup
```

## Usage

Import the package
```go
import "github.com/h2non/findup"
```

Find a file
```go
path, err := findup('findup.go')
fmt.Println(path) // -> /full/path/to/findup.go
```

Find a file using a Glob pattern
```go
path, err := findup('findup.*')
fmt.Println(path) // -> /full/path/to/findup.go
```

See also [GoDoc](https://godoc.org/github.com/h2non/findup)

## Development

```bash
go test
```

## License 

MIT - Tomas Aparicio
