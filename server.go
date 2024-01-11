package main

import (
	"log"
	"net/http"

	"git.jmbit.de/jmb/www-jmbit-de/public"
)

func main() {
	// Register a custom handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Specify the file path you want to block
		blockFilePath := "/public.go"

		// Check if the requested path matches the blocked file path
		if r.URL.Path == blockFilePath {
			// Return a 404 Not Found error
			http.NotFound(w, r)
			return
		}

		// For other paths, serve the files using the file server
		http.FileServer(http.FS(public.HtmlFS)).ServeHTTP(w, r)
	})

	// Start the HTTP server on port 80
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
