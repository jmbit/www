package web

import (
  "net/http"
  "log"
)


func ToolsWebHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")

  component := Tools()
    err := component.Render(r.Context(), w)
    if err != nil {
          http.Error(w, err.Error(), http.StatusBadRequest)
          log.Printf("Error rendering in ToolsWebHandler: %e", err)
  }
}
