package main

import (
	"git.jmbit.de/jmb/www-jmbit-de/internal/config"
	"git.jmbit.de/jmb/www-jmbit-de/internal/server"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	server := server.NewServer(config.Host, config.Port, config.Debug)
	server.ListenAndServe()
}
