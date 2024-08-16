package html

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

var (
	//go:embed build/*
	html embed.FS

	docs fs.FS

	pathList []string
)

func init() {
	var err error
	docs, err = fs.Sub(html, "build")
	if err != nil {
		panic(err)
	}
	pathList, err = traverseFileSystem()
	if err != nil {
		panic(err)
	}
}

type (
	dependencies interface{}
)

func NewHandler(r dependencies) http.Handler {
	return http.FileServerFS(docs)
}

func PathList() []string {
	return pathList
}

func traverseFileSystem() ([]string, error) {
	var paths []string
	err := fs.WalkDir(docs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		paths = append(paths, "/"+path)
		if strings.HasSuffix(path, "index.html") {
			paths = append(paths, "/"+strings.TrimSuffix(path, "index.html"))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paths, nil
}
