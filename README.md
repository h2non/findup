# findup.go

Find the first file matching in a given directory or the nearest ancestor directory up to root

> Work in progress!

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
path, err := findup('config.json')
```

## License 

MIT - Tomas Aparicio
