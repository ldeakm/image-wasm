package main

import (
	"net/http"
	"os"
)

func main() {
	// Default to serve index
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		contentHandler("index.html", w, r)
	})

	// Serve any dependancy files
	http.HandleFunc("/static/wasm_exec.js", func(w http.ResponseWriter, r *http.Request) {

		contentHandler("static/wasm_exec.js", w, r)
	})

	// Serve the wasm file and set headers
	http.HandleFunc("/app.wasm", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/wasm")
		contentHandler("app.wasm", w, r)

	})

	http.ListenAndServe(":8080", nil)
}

func contentHandler(path string, w http.ResponseWriter, r *http.Request) {

	item, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fileStat, err := item.Stat()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	http.ServeContent(w, r, path, fileStat.ModTime(), item)

	return
}
