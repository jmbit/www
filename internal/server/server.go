package server

import (
	"fmt"
	"git.jmbit.de/jmb/www-jmbit-de/internal/middlewares"
	"log"
	"net/http"
	"time"
)

func NewServer(host string, port int, logging bool) *http.Server {

	middlewareStack := middlewares.CreateStack()
	if logging {
		log.Println("Enabling HTTP logging")
		middlewareStack = middlewares.CreateStack(
			middlewares.Logging,
			middlewares.TerryPratchet,
		)
	} else {
		middlewareStack = middlewares.CreateStack(
			middlewares.TerryPratchet,
		)
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      middlewareStack(RegisterRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return server
}
