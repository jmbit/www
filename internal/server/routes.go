package server

import (
	"net/http"

	"git.jmbit.de/jmb/www-jmbit-de/internal/web"
)

// RegisterRoutes registers all the routes
func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
  mux.HandleFunc("/tools/", web.ToolsWebHandler)
	mux.HandleFunc("/", web.HugoWebHandler)

	return mux
}
