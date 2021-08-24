package git

import (
	"encoding/json"
	"io"
	"net/http"
)

type File struct {
	Filename, Status, Contents_url string
}

func FetchFiles(url string) []File {
	return DecodeFiles(Fetch(url))
}

func Fetch(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.Body
}

func DecodeFiles(r io.Reader) []File {
	decoder := json.NewDecoder(r)
	var files []File
	err := decoder.Decode(&files)
	if err != nil {
		panic(err)
	}
	return files
}
