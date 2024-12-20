package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host string
	Port int
  Debug bool
}

func GetConfig() (*Config, error) {
	config := &Config{}

	if os.Getenv("HOST") != "" {
		config.Host = os.Getenv("HOST")
	} else {
		config.Host = "0.0.0.0"
	}

	if os.Getenv("PORT") != "" {
		port := os.Getenv("PORT")
		var err error
		config.Port, err = strconv.Atoi(port)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		config.Port = 80
	}
  
  if os.Getenv("DEBUG") != "" {
    debug := os.Getenv("DEBUG")
    var err error
    config.Debug, err = strconv.ParseBool(debug)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    config.Debug = false
  }

	return config, nil
}
