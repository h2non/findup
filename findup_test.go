package findup

import (
	"os"
	"path"
	"path/filepath"
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
	testFindFilePath(t, ".tmp/a")
	testFindFilePath(t, ".tmp/a/b")
	testFindFilePath(t, ".tmp/a/b/c")
}

func testFindFilePath(t *testing.T, dirPath string) {
	cwd, _ := os.Getwd()
	goodPath := getFullPath("findup.go")

	tmpDir := filepath.Join(cwd, dirPath)
	os.MkdirAll(tmpDir, 0666)

	path, _ := lookupFile(tmpDir, "findup.go")
	if path != goodPath {
		t.Fatalf("file was not found")
	}

	os.RemoveAll(".tmp")
}

func getFullPath(filename string) string {
	cwd, _ := os.Getwd()
	return path.Join(cwd, filename)
}
