package web

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed templates/*.gohtml
var Content embed.FS

//go:embed static/*
var Static embed.FS

var fileServer = http.FileServer(http.FS(Static)).ServeHTTP

type entryInfo struct {
	sha string
}

var entries map[string]entryInfo

func init() {
	entries = make(map[string]entryInfo)
	e, err := Static.ReadDir("static")
	if err != nil {
		panic(err)
	}
	for _, f := range e {
		if f.IsDir() {
			continue
		}
		name := f.Name()
		fileBytes, err := Static.ReadFile("static/" + name)
		if err != nil {
			panic(err)
		}
		sha := sha256.New()
		sha.Write(fileBytes)
		entries[name] = entryInfo{
			sha: hex.EncodeToString(sha.Sum(nil)),
		}
	}
}

var ServeStatic = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	file := mux.Vars(r)["file"]
	if file == "" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	entry, ok := entries[file]
	if !ok {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	if r.Header.Get("If-None-Match") == entry.sha {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("ETag", entry.sha)
	w.Header().Set("Cache-Control", "public, max-age=3600")

	fileServer(w, r)
})
