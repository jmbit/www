package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public/*
var publicFS embed.FS


func main() {
	// Register a custom handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

    fsroot, err := fs.Sub(publicFS, "public")
    if err != nil {
            panic(err)
    }
    http.FileServer(http.FS(fsroot)).ServeHTTP(w,r)

	})

	// Start the HTTP server on port 80
	err := http.ListenAndServe(":1313", nil)
	if err != nil {
		log.Fatal(err)
	}
}
