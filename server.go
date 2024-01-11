package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed public/*
var HtmlFS embed.FS

func main() {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Recovery())

	router.StaticFS("/", http.FS(HtmlFS))
	err := router.Run("0.0.0.0:80")
	if err != nil {
		log.Fatal(err)
	}
}
