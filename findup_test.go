package findup

import (
	"testing"
)

func FindInCurrentDirectory(t *testing.T) {
	path, error := Find(".gitignore")
	Equal(t, path, "pepe")
}
