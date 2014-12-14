package findup

import (
	"os"
	"path"
	"testing"
)

func TestExistentFileInTheCurrentDirectory(t *testing.T) {
	path, _ := Find(".gitignore")
	if path != getFullPath(".gitignore") {
		t.Fatalf("file not found")
	}
}

func TestCannotFindTheFileInTheCurrentDirectory(t *testing.T) {
	path, _ := Find("not-found")
	if path != "" {
		t.Fatalf("file was found")
	}
}

func TestFileWithGlobPatternInTheCurrentDirectory(t *testing.T) {
	path, _ := Find("findup.*")
	if path != getFullPath("findup.go") {
		t.Fatalf("file not found")
	}
}

func TestCannotFindGlobFileInTheCurrentDirectory(t *testing.T) {
	path, _ := Find("not-*")
	if path != "" {
		t.Fatalf("file was found")
	}
}

func TestFindFileInNearestDirectories(t *testing.T) {
	tmpDir := ".tmp/a/b/c"
	os.MkdirAll(tmpDir, 0666)
	os.Chdir(tmpDir)

	path, _ := Find("findup.go")
	if path != getFullPath("findup.go") {
		t.Fatalf("file was found")
	}

	os.RemoveAll(".tmp")
	os.Chdir("../../../")
}

func getFullPath(filename string) string {
	cwd, _ := os.Getwd()
	return path.Join(cwd, filename)
}
