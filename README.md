# findup.go

A [Go](http://golang.org) tiny package to **find the first** file matching in a given directory or the nearest ancestor directory up to root, with additional Glob pattern support

## Installation

Using `go get`
```bash
$ go get github.com/h2non/findup
```

## API

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
$ go test
```

## License 

MIT - Tomas Aparicio
