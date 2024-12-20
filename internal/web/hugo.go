package web

import (
	"io/fs"
	"net/http"

	"git.jmbit.de/jmb/www-jmbit-de/hugo"
)

func HugoWebHandler(w http.ResponseWriter, r *http.Request) {
	fsroot, err := fs.Sub(hugo.PublicFS, "public")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Cache-Control", "max-age=3600")
	http.FileServer(http.FS(fsroot)).ServeHTTP(w, r)
}
