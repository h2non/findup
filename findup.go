package findup

import (
	"os"
	"path/filepath"
)

func Find(filename string) (string, error) {
	matches, err := filepath.Find(os.Getcwd() + filename)
	//fmt.Printf(matches)
	return matches, err
}
