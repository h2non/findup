# findup.go [![Build Status](https://travis-ci.org/h2non/findup.png)](https://travis-ci.org/h2non/findup) [![GitHub release](http://img.shields.io/github/release/h2non/findup.svg)]()

A [Go](http://golang.org) tiny package to **find the first file matching** in the current directory or the nearest ancestor directory up to root, with additional [Glob patterns](http://en.wikipedia.org/wiki/Glob_%28programming%29) support

It works in Go >= 1.1

## Installation

```bash
go get github.com/h2non/findup
```

## Usage

See also [Godoc](https://godoc.org/github.com/h2non/findup)

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

## Development

```bash
go test
```

## License 

MIT - Tomas Aparicio
