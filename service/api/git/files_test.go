package git

import (
	"strings"
	"testing"
)

func TestDecodeFiles(t *testing.T) {
	const jsonFiles = `
	[
		{
			"filename": "doc.go",
			"status": "removed",
			"contents_url": "https://api.github.com/repos/owner/repo/contents/myFile.txt"
		}
	]
	`

	files := DecodeFiles(strings.NewReader(jsonFiles))

	assertStringEquals(t, files[0].Filename, "doc.go")
	assertStringEquals(t, files[0].Status, "removed")
	assertStringEquals(
		t, files[0].Contents_url,
		"https://api.github.com/repos/owner/repo/contents/myFile.txt")
}

// helpers

func assertStringEquals(t *testing.T, s string, s2 string) bool {
	ret := s == s2

	if s != s2 {
		t.Fatalf("Want %v Got %v", s2, s)
	}

	return ret
}
