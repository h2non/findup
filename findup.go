package findup

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

func Find(filename string) (string, error) {
	cwd, _ := os.Getwd()
	return lookupFile(cwd, filename)
}

func lookupFile(basepath string, filename string) (string, error) {
	matches, err := filepath.Glob(path.Join(basepath, filename))
	if len(matches) == 0 {
		return lookupInNearestDir(basepath, filename)
	}
	return matches[0], err
}

func lookupInNearestDir(basepath string, filename string) (string, error) {
	if basepath == "/" {
		return "", errors.New("file not found")
	}
	nearest := path.Dir(basepath)
	return lookupFile(nearest, filename)
}
